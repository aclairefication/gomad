package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/castillobgr/sententia"
)

type ReturnedWords struct {
	Nouns        []string `json:"noun"`
	PresentVerbs []string `json:"verb_present"`
	Verbs        []string `json:"verb"`
	Adjectives   []string `json:"adjective"`
}

func callwordblanks() {
	//Word Blanks Mad Libs generator

	//Hardcoding is so boring
	// partsofspeech := "&posArray%5B%5D=noun&posArray%5B%5D=verb_present&posArray%5B%5D=verb&posArray%5B%5D=adjective"
	// wordblanksurl := "https://www.wordblanks.com/scripts/wb_ajax.php?token=diZo19~pw9wzzTmjssdNBAcJ2SQECyG8Ip.LwefrOnM-&method=getRandomWordSet&posArray%5B%5D=noun&posArray%5B%5D=verb_present&posArray%5B%5D=verb&posArray%5B%5D=adjective"

	wordblanksapibase := "https://www.wordblanks.com/scripts/wb_ajax.php?"
	token := "diZo19~pw9wzzTmjssdNBAcJ2SQECyG8Ip.LwefrOnM-"
	//Getting words
	method := "getRandomWordSet"

	// Want an slice of parts of speech and then range across to build this URL
	parts_of_speech := []string{"noun", "verb_present", "verb", "adjective"}
	var speechparams string
	for _, part := range parts_of_speech {
		speechparams = speechparams + "&posArray%5B%5D=" + part
	}
	// fmt.Println("Parts of speech are", speechparams)

	fmt.Println("Calling Word Blanks")

	// wordblanksurl := wordblanksapibase + "token=" + token + "&" + "method=" + method + partsofspeech
	wordblanksurl := wordblanksapibase + "token=" + token + "&" + "method=" + method + speechparams
	fmt.Println("Request to URL ", wordblanksurl)

	//Modified code from https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9

	var record ReturnedWords

	// Build the request
	req, err := http.NewRequest("GET", wordblanksurl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Nouns list = ", record.Nouns)
	fmt.Println("Present tense verbs = ", record.PresentVerbs)
	fmt.Println("Verbs = ", record.Verbs)
	fmt.Println("Adjectives = ", record.Adjectives)

}

func callsententia() {
	fmt.Println("Calling Sententia")

	//Using sententia to generate Mad Libs with adjectives & nouns
	sentence, err := sententia.Make("Aw yis, {{ adjective }} {{ nouns }}.")
	if err != nil {
		panic(err)
	}
	fmt.Println(sentence)

}

func main() {
	callwordblanks()
	callsententia()
}
