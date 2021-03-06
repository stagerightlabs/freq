package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
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

func TestLetters(t *testing.T) {
	sampleText := "GFEDCBAGFEDCBA"
	expectedLetters := []string{"A", "B", "C", "D", "E", "F", "G"}
	ls := NewLetterSet()
	CountLetters(sampleText, &ls)

	if !reflect.DeepEqual(ls.Letters(), expectedLetters) {
		t.Errorf("Expected Letters() to return %v", expectedLetters)
	}
}

func TestMostCommonLetters(t *testing.T) {
	sampleText := "AAAABBBCCD"
	expectedLetters := []string{"A"}
	ls := NewLetterSet()
	CountLetters(sampleText, &ls)

	if !reflect.DeepEqual(ls.MostCommonLetters(), expectedLetters) {
		t.Errorf("Expected Letters() to return %v", expectedLetters)
	}
}

func TestApiFreqHandler(t *testing.T) {

	// Simulate our POST form data
	form := url.Values{}
	form.Add("text", "GFEDCBAGFEDCBA")

	// Simulate a request and attach the form to it
	req, err := http.NewRequest("POST", "/freq", strings.NewReader(form.Encode()))
	req.Form = form
	if err != nil {
		log.Fatal(err.Error())
	}

	// Simulate a response to receive the handler output
	resp := httptest.NewRecorder()

	// Call the handler we are testing
	apiFreqHandler(resp, req)

	// Check the response body
	if resp.Body.String() != `{"text":"GFEDCBAGFEDCBA","counts":{"A":2,"B":2,"C":2,"D":2,"E":2,"F":2,"G":2},"totalLetters":7,"mostFrequent":["A","B","C","D","E","F","G"]}` {
		t.Error("Received an invalid API response from /freq")
	}
}
