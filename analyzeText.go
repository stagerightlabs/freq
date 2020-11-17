package main

// analyzeText performs frequency analysis on a text string
func analyzeText(text string) {
	ls := NewLetterSet()
	ls.text = text
	CountLetters(text, &ls)
	ls.Print()
}
