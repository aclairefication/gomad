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
