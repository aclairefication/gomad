package main

import (
	"testing"
)

//Generate sentence should return a non-empty string
func TestGenerateNonEmptySentence(t *testing.T) {
	sentence := generateSentence()
	if sentence == "" {
		t.Errorf("generateSentence was incorrect, got empty string but want not empty string")
	}
}
