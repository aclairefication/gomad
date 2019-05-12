package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
  "io/ioutil"
  "strconv"

  "github.com/pact-foundation/pact-go/dsl"
  "github.com/pact-foundation/pact-go/types"
)

//directly from pact-go\dsl\mock_service_test.go
// Simple mock server for testing. This is getting confusing...
func setupMockServer(success bool, t *testing.T) *httptest.Server {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    _, err := ioutil.ReadAll(r.Body)
    r.Body.Close()
    if err != nil {
      log.Fatal(err)
    }

    if success {
      fmt.Fprintln(w, "Hello, client")
    } else {
      http.Error(w, "something went wrong\n", http.StatusInternalServerError)
    }
  }))

  return ts
}

//Expects a URL but doesn't have error-handling
func getPort(input string) string {
  splits := strings.SplitAfter(input, ":")
  port := splits[2]
  // fmt.Println("split to get " + port)
  return port
}

//Adapted from pact-go/examples/consumer_test.go
func TestConsumer(t *testing.T) {

  type ResponseWords struct {
    Adjective   string `json:"adjective" pact:"example=better,regex=^better$"`
    Noun        string `json:"noun" pact:"example=software,regex=^software$"`
    Verb        string `json:"verb" pact:"example=do,regex=^do$"`
  }

  ms := setupMockServer(true, t)
  defer ms.Close()

  // fmt.Println("mock server URL is " + ms.URL)

  port, _ := strconv.Atoi(getPort(ms.URL))
  // port, err := strconv.Atoi(getPort(ms.URL))
  // if err == nil {
  //     fmt.Println(port)
  // }
  // fmt.Println("port should be " + getPort(ms.URL))

	// Create Pact client
	pact := &dsl.Pact{
    Server: &types.MockServer{
      Port: port,
    },
		Consumer: "MadLibsTemplate",
		Provider: "WordChooser",
    Host:     "localhost",
	}
	// Shuts down Mock Service when done
  defer pact.Teardown()

  // Pass in test case
  var test = func() error {
    u := fmt.Sprintf("http://localhost:%d/words", pact.Server.Port)
    req, err := http.NewRequest("GET", u, strings.NewReader(`{"adjective":"better"}`))

    if err != nil {
      return err
    }
    if _, err = http.DefaultClient.Do(req); err != nil {
      return err
    }

    return err
  }

  // Contract will be GET with parts of speech and number of instances needed
  pact.
    AddInteraction().
    Given("Parts of speech for Agile Manifesto's first sentence exist").
    UponReceiving("A request to get words").
    WithRequest(dsl.Request{
      Method:  "GET",
      Path:    "/words",
      Body: map[string]int{
        "adjective": 1,
        "noun": 1,
        "verb": 1,
      },
    }).
    WillRespondWith(dsl.Response{
      Status:  200,
      Body:    dsl.Match(&ResponseWords{}),
    })

  // Verify
  if err := pact.Verify(test); err != nil {
    log.Fatalf("Error on Verify: %v", err)
  }

  fmt.Println("Test Passed!")
}