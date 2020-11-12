package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"unicode"
)

func main() {

	var text string
	var file string

	// Read flags
	flag.StringVar(&text, "t", "", "Specify text to analyze")
	flag.StringVar(&text, "text", "", "Specify text to analyze")
	flag.StringVar(&file, "f", "", "Specify a file to read")
	flag.StringVar(&file, "file", "", "Specify a file to read")
	flag.Parse()

	// Use the flag contents to determine our handler method
	if len(file) > 0 {
		AnalyzeFile(file)
	} else if len(text) > 0 {
		AnalyzeText(text)
	} else {
		fmt.Println("You did not specify any input.")
	}
}

// AnalyzeText performs frequency analysis on a text string
func AnalyzeText(text string) {
	ls := NewLetterSet()
	ls.text = text
	CountLetters(text, &ls)
	ls.Print()
}

// AnalyzeFile performs frequency analysis on the contents of a file
func AnalyzeFile(file string) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ls := NewLetterSet()
	ls.file = file
	CountLetters(string(contents), &ls)
	ls.Print()
}

// CountLetters returns a LetterSet containing a frequency
// count for every letter that appears in the source text
func CountLetters(text string, ls *LetterSet) {

	var char rune

	for _, c := range text {
		char = unicode.ToUpper(c)

		if char >= 'A' && char <= 'Z' {
			ls.counts[char] = ls.counts[char] + 1
			ls.total++
		}
	}
}

// LetterSet represents a frequency count for a group of letters
type LetterSet struct {
	counts map[rune]int
	total  int64
	text   string
	file   string
}

// NewLetterSet returns an empty LetterSet
func NewLetterSet() LetterSet {
	return LetterSet{counts: make(map[rune]int), total: 0}
}

// Print a LetterSet to the console
func (l LetterSet) Print() {

	if len(l.file) > 0 {
		fmt.Println("File:", l.file)
	}

	fmt.Println("Frequency Analysis:")

	// Prepare to traverse our counts map alphabetically
	keys := make([]string, 0, len(l.counts))
	for r := range l.counts {
		keys = append(keys, string(r))
	}
	sort.Strings(keys)

	// Display our letter counts in alphabeticall order
	for _, v := range keys {
		letter := []rune(v)[0]
		count := l.counts[letter]
		percentage := float64(count) / float64(l.total)
		fmt.Printf("%c: %v (%.2f%%)\n", letter, count, percentage)
	}

	// Display the total letter count
	fmt.Printf("Total Letters: %v\n", l.total)
}

// Empty indicates wether or not the letter set has any contents
func (l LetterSet) Empty() bool {
	return len(l.counts) == 0
}
