package main

import (
	"fmt"
	"testing"
)

func TestEmptyLetterSet(t *testing.T) {
	ls := NewLetterSet()

	if !ls.Empty() {
		t.Error("Expected LetterSet to be empty.")
	}
}

func TestCountLetters_Alphabetic(t *testing.T) {
	sampleText := "ABCDEFGABCDEFG"
	ls := NewLetterSet()
	CountLetters(sampleText, &ls)

	if ls.Empty() {
		t.Error("Expected LetterSet to have contents")
	}

	for _, letter := range sampleText {
		if ls.counts[letter] != 2 {
			t.Errorf("Expected letter %c to have a frequency count of 1", letter)
		}
	}

	if ls.total != int64(len(sampleText)) {
		expectedLength := fmt.Sprint(len(sampleText))
		t.Errorf("Expected LetterSet to have a total of %v", expectedLength)
	}
}

func TestCountLetters_NonAlphabetic(t *testing.T) {
	sampleText := "123456789.,!@#$%&*();?"
	ls := NewLetterSet()
	CountLetters(sampleText, &ls)

	if !ls.Empty() {
		t.Error("Expected LetterSet to be empty")
	}
}
