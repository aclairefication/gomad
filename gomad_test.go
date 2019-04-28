package main

import (
	"strings"
	"testing"
)

//Generate sentence should return a non-empty string
func TestGeneratedSentenceIsNotEmpty(t *testing.T) {
	sentence := generateSentence()
	if sentence == "" {
		t.Error("generateSentence was incorrect, got empty string but want not empty string")
	}
}

//Generate sentence should return a string that contains "Aw yis, "
func TestGeneratedSentenceContainsAwYis(t *testing.T) {
	sentence := generateSentence()

	// if sentence != "Aw yis, " {
	if !strings.Contains(sentence, "Aw yis, ") {
		t.Error("generateSentence was incorrect. Should have contained 'Aw yis, ' but was '" + sentence + "'")
	}
}

//Generate sentence should return a string starting with "Aw yis, "
func TestGeneratedSentenceStartsWithAwYis(t *testing.T) {
	sentence := generateSentence()

	if !strings.HasPrefix(sentence, "Aw yis, ") {
		t.Error("generateSentence was incorrect. Should have started with 'Aw yis, ' but was '" + sentence + "'")
	}
}

//Generate sentence should return a sentence ending with a period
func TestGeneratedSentenceEndsWithAPeriod(t *testing.T) {
	sentence := generateSentence()

	if !strings.HasSuffix(sentence, ".") {
		t.Error("generateSentence was incorrect. Should have ended with a period but was '" + sentence + "'")
	}
}

//Generated sentence should contain 4 words
func TestGeneratedSentenceContainsFourWords(t *testing.T) {
	sentence := generateSentence()
	spaceCount := strings.Count(sentence, " ")

	if spaceCount != 3 {
		t.Error("generateSentence was incorrect. Should have contained 3 spaces (i.e. 4 words) but was '" + sentence + "'")
	}
}

func TestFillSentence(t *testing.T) {
	sentencePattern := "Aw yis, {{ adjective }} {{ nouns }}."
	sentencePrefix := strings.Split(sentencePattern, "{{")[0]
	t.Log("Prefix is '" + sentencePrefix + "'")
	patternsInString := strings.Split(sentencePattern, "}}")
	numPatterns := len(patternsInString)
	sentenceSuffix := strings.Split(sentencePattern, "}}")[numPatterns-1]
	t.Log("Suffix is '" + sentenceSuffix + "'")

	filledSentence := fillSentence(sentencePattern)

	if !strings.HasPrefix(filledSentence, sentencePrefix) && !strings.HasSuffix(filledSentence, sentenceSuffix) {
		t.Error("generateSentence was incorrect. Should have started with '" + sentencePrefix + "' and end with '" + sentenceSuffix + "' but was '" + filledSentence + "'")
	}

}

func TestFillManySentences(t *testing.T) {
	var sentencePatterns = []struct {
		pattern string
		prefix  string
		suffix  string
	}{
		{"Aw yis, {{ adjective }} {{ nouns }}.", "Aw yis, ", "."},
		{"Hark! {{ adjective }} {{ nouns }}!", "Hark! ", "!"},
		{"{{ adjective }} {{ nouns }} are at the door?", "", " are at the door?"},
	}

	for _, thisSentence := range sentencePatterns {
		filledSentence := fillSentence(thisSentence.pattern)
		t.Log("Resulting sentence: " + filledSentence)

		if !strings.HasPrefix(filledSentence, thisSentence.prefix) && !strings.HasSuffix(filledSentence, thisSentence.suffix) {
			t.Error("generateSentence was incorrect. Should have started with '" + thisSentence.prefix + "' and end with '" + thisSentence.suffix + "' but was '" + filledSentence + "'")
		}
	}

}
