package main

import (
	"flag"
	"fmt"
	"sort"
	"unicode"
)

var text string

func init() {
	flag.StringVar(&text, "t", "", "Specify text to analyze")
	flag.StringVar(&text, "text", "", "Specify text to analyze")
}

func main() {
	flag.Parse()
	set := countLetters(text)
	set.print()
}

// Create a new frequency count for a given string
func countLetters(text string) letterSet {

	ls := newLetterSet()
	var char rune

	for _, c := range text {
		char = unicode.ToUpper(c)

		if char >= 'A' && char <= 'Z' {
			ls.counts[char] = ls.counts[char] + 1
			ls.total++
		}
	}

	return ls
}

// A letter set represents a completed frequency analysis
type letterSet struct {
	counts map[rune]int
	total  int64
}

func newLetterSet() letterSet {
	return letterSet{counts: make(map[rune]int), total: 0}
}

// Print a letterSet to the console
func (l letterSet) print() {

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
	fmt.Printf("Total: %v\n", l.total)
}
