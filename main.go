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
	ls := CountLetters(text)
	ls.Print()
}

// CountLetters returns a LetterSet containing a frequency
// count for every letter that appears in the source text
func CountLetters(text string) LetterSet {

	ls := NewLetterSet()
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

// LetterSet represents a frequency count for a group of letters
type LetterSet struct {
	counts map[rune]int
	total  int64
}

// NewLetterSet returns an empty LetterSet
func NewLetterSet() LetterSet {
	return LetterSet{counts: make(map[rune]int), total: 0}
}

// Print a LetterSet to the console
func (l LetterSet) Print() {

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

// Empty indicates wether or not the letter set has any contents
func (l LetterSet) Empty() bool {
	return len(l.counts) == 0
}
