package main

import (
	"io/ioutil"
)

// analyzeFile performs frequency analysis on the contents of a file
func analyzeFile(file string) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	ls := NewLetterSet()
	ls.file = file
	CountLetters(string(contents), &ls)
	ls.Print()

	return nil
}
