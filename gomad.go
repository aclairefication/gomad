package main

import (
        "fmt"
        "strings"

        "github.com/castillobgr/sententia"
)

func generateSentence() string {
        // fmt.Println("Calling generateSentence")
        //Using sententia to generate Mad Libs with adjectives & nouns
        sentence, err := sententia.Make("Aw yis, {{ adjective }} {{ nouns }}.")
        if err != nil {
                panic(err)
        }
        return sentence
}

func fillSentence(sentenceWithBlanks string) string {
        fmt.Println("Calling fillSentence with '" + sentenceWithBlanks + "'")
        sentence, err := sententia.Make(sentenceWithBlanks)
        if err != nil {
                panic(err)
        }
        fmt.Println("fillSentence created '" + sentence + "'")
        return sentence
}

func main() {
        //Function with hard-coded string pattern and no return
        callWordBlanks()

        //Function with hard-coded string pattern and result return
        madlib := generateSentence()
        fmt.Println("madlib sentence is " + madlib)

        //Function with dynamic input and result return
        sentenceToFill := "Aw yis, {{ adjective }} {{ nouns }}."
        completedSentence := fillSentence(sentenceToFill)
        fmt.Println("another madlib sentence: " + completedSentence)

        //Function with dynamic input and result return
        sentenceToFill2 := "Oh look, {{ an adjective }} {{ noun }}!"
        completedSentence2 := fillSentence(sentenceToFill2)
        fmt.Println("another madlib sentence: " + completedSentence2)

        //Function with dynamic input and result return that uses Actions from Go's text/template package
        //Define new action
        customActions := map[string]interface{}{
                "capitalize": strings.Title,
        }
        sententia.AddActions(customActions)
        //Use new action
        completedSentence3, err := sententia.Make(
                "She wrote a book called '{{ capitalize (an adjective) }} {{ capitalize noun }}'",
        )
        if err != nil {
                panic(err)
        } else {
                fmt.Println("actions madlib: " + completedSentence3)
        }

}