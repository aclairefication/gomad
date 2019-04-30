package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReturnedWords struct {
	Nouns        []string `json:"noun"`
	PresentVerbs []string `json:"verb_present"`
	Verbs        []string `json:"verb"`
	Adjectives   []string `json:"adjective"`
}

func buildwordBlanksURL() string {

	token := "2UdLWTsozqnPNltOB4n2h6X8nf7lAH3m"

	//Hardcoding is so boring
	// URL for GET from https://www.wordblanks.com/mad-libs/computing/moss-memory/
	// wordBlanksURL := "https://www.wordblanks.com/scripts/wb_ajax.php?token=diZo19~pw9wzzTmjssdNBAcJ2SQECyG8Ip.LwefrOnM-&method=getRandomWordSet&posArray%5B%5D=noun&posArray%5B%5D=verb_present&posArray%5B%5D=verb&posArray%5B%5D=adjective"
	// partsofspeech := "&posArray%5B%5D=noun&posArray%5B%5D=verb_present&posArray%5B%5D=verb&posArray%5B%5D=adjective"
	// wordBlanksURL := wordblanksapibase + "token=" + token + "&" + "method=" + method + partsofspeech

	wordblanksapibase := "https://www.wordblanks.com/scripts/wb_ajax.php?"
	method := "getRandomWordSet"

	// Want a slice of parts of speech and then range across to build this URL
	partsOfSpeech := []string{"noun", "verb_present", "verb", "adjective"}
	var speechparams string
	for _, part := range partsOfSpeech {
		speechparams = speechparams + "&posArray%5B%5D=" + part
	}

	fmt.Println("Calling Word Blanks")

	wordBlanksURL := wordblanksapibase + "token=" + token + "&" + "method=" + method + speechparams
	fmt.Println("Request to URL ", wordBlanksURL)

	return wordBlanksURL
}

//Modified code from https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9

func buildRequest(requestMethod string, wordBlanksURL string) *http.Request {
	// Build the request
	request, err := http.NewRequest("GET", wordBlanksURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	return request
}

func sendRequest(request *http.Request, client http.Client) *http.Response {
	// Send the request via a client
	// Do sends an HTTP request and returns an HTTP response
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	return response
}

func callWordBlanks() {

	wordBlanksURL := buildwordBlanksURL()

	var madLibWords ReturnedWords

	request := buildRequest("GET", wordBlanksURL)

	// For control over HTTP client headers, redirect policy, and other settings, create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	response := sendRequest(request, *client)

	// Callers should close response.Body when done reading from it
	// Defer the closing of the body
	defer response.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&madLibWords); err != nil {
		log.Println(err)
	}

	fmt.Println("Nouns list = ", madLibWords.Nouns)
	fmt.Println("Present tense verbs = ", madLibWords.PresentVerbs)
	fmt.Println("Verbs = ", madLibWords.Verbs)
	fmt.Println("Adjectives = ", madLibWords.Adjectives)

}
