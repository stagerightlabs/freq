package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"unicode"
)

// LetterSet represents a frequency count for a group of letters
type LetterSet struct {
	counts map[rune]int
	total  int64
	text   string
	file   string `json:"-"`
}

// NewLetterSet returns an new LetterSet
func NewLetterSet() LetterSet {
	return LetterSet{
		counts: make(map[rune]int),
		total:  0,
	}
}

// Print a LetterSet to the console
func (l LetterSet) Print() {
	if len(l.file) > 0 {
		fmt.Println("File:", l.file)
	}

	fmt.Println("Frequency Analysis:")

	// Display our letter counts in alphabetical order
	for _, v := range l.Letters() {
		letter := []rune(v)[0]
		count := l.counts[letter]
		percentage := float64(count) / float64(l.total)
		fmt.Printf("%c: %v (%.2f%%)\n", letter, count, percentage)
	}

	// Display the total letter count
	fmt.Printf("Total Letters: %v\n", l.total)
	fmt.Printf("Most frequent: %v\n", l.MostCommonLetters())
}

// Letters returns a slice of strings containing the letters
// in the LetterSet, sorted alphabetically
func (l LetterSet) Letters() []string {
	// Prepare to traverse our counts map alphabetically
	keys := make([]string, 0, len(l.counts))
	for r := range l.counts {
		keys = append(keys, string(r))
	}
	sort.Strings(keys)

	return keys
}

// Empty indicates wether or not the letter set has any contents
func (l *LetterSet) Empty() bool {
	return len(l.counts) == 0
}

// MarshalJSON converts a LetterSet to JSON
func (l *LetterSet) MarshalJSON() ([]byte, error) {

	counts := make(map[string]int)
	total := 0

	for r := range l.counts {
		counts[string(r)] = l.counts[r]
		total++
	}

	return json.Marshal(&struct {
		Text         string         `json:"text"`
		Counts       map[string]int `json:"counts"`
		TotalLetters int            `json:"totalLetters"`
		MostFrequent []string       `json:"mostFrequent"`
	}{
		Counts:       counts,
		Text:         l.text,
		TotalLetters: total,
		MostFrequent: l.MostCommonLetters(),
	})
}

// MostCommonLetters returns a slice of strings containing
// the letters with the highest frequency count in the
// LetterSet, sorted alphabetically
func (l *LetterSet) MostCommonLetters() []string {
	mostFrequent := []string{}
	max := 0

	// Find the highest frequency count
	for _, count := range l.counts {
		if count > max {
			max = count
		}
	}

	// Retrieve the letters with the highest count
	for rune, count := range l.counts {
		if count == max {
			mostFrequent = append(mostFrequent, string(rune))
		}
	}

	sort.Strings(mostFrequent)

	return mostFrequent
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
