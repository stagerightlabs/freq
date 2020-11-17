package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handle an API request
func apiFreqHandler(w http.ResponseWriter, r *http.Request) {
	ls := NewLetterSet()
	CountLetters(r.FormValue("text"), &ls)
	ls.text = r.FormValue("text")

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := ls.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(jsonData))
}
