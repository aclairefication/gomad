package main_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//HTTP request to wordblanks
/*
	Nouns        []string `json:"noun"`
	PresentVerbs []string `json:"verb_present"`
	Verbs        []string `json:"verb"`
	Adjectives   []string `json:"adjective"`

	fmt.Println("Nouns list = ", record.Nouns)
	fmt.Println("Present tense verbs = ", record.PresentVerbs)
	fmt.Println("Verbs = ", record.Verbs)
	fmt.Println("Adjectives = ", record.Adjectives)
*/

//Modeled after https://tutorialedge.net/golang/advanced-go-testing-tutorial/#mocking-http-requests
func TestHttpGetRequestForWordBlanks(t *testing.T) {
	token := "2UdLWTsozqnPNltOB4n2h6X8nf7lAH3m"

	sampleNouns := "[\"mission\",\"anchor\",\"yogurt\",\"amusement park\"]"
	samplePresentVerbs := "[\"inspiring\",\"wobbling\",\"mailing\"]"
	sampleVerbs := "[\"crack\",\"breathe\",\"climb\",\"sip\",\"make\"]"
	sampleAdjectives := "[\"vociferous\",\"chatty\",\"rigorous\",\"scraggly\",\"strategic\",\"hurried\"]"
	sampleResponse := "{\"noun\":[" + sampleNouns + ", \"verb_present\":[" + samplePresentVerbs + ", \"verb\":" + sampleVerbs + ", \"adjective\":" + sampleAdjectives + "}"

	targetURL := "https://www.wordblanks.com/scripts/wb_ajax.php?token=" + token + "&method=getRandomWordSet&posArray%5B%5D=noun&posArray%5B%5D=verb_present&posArray%5B%5D=verb&posArray%5B%5D=adjective"

	handler := func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, sampleResponse)
	}

	//Mock response
	request := httptest.NewRequest("GET", targetURL, nil)
	writer := httptest.NewRecorder()
	handler(writer, request)

	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	bodyS := string(body)
	fmt.Println(bodyS)
	if 200 != response.StatusCode {
		t.Fatal("Status code is " + string(response.StatusCode) + " not 200 OK.")
	}
	if bodyS != sampleResponse {
		t.Fatal("Body content is not expected: " + bodyS )
	}

}
