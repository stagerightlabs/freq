package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"unicode"
)

func main() {

	var text string
	var file string
	var serve bool
	var host string
	var port string

	// Read flags
	flag.StringVar(&text, "t", "", "Specify text to analyze")
	flag.StringVar(&text, "text", "", "Specify text to analyze")
	flag.StringVar(&file, "f", "", "Specify a file to read")
	flag.StringVar(&file, "file", "", "Specify a file to read")
	flag.BoolVar(&serve, "s", false, "Launch the web server")
	flag.BoolVar(&serve, "serve", false, "Launch the web server")
	flag.StringVar(&host, "h", "0.0.0.0", "Specify the server host")
	flag.StringVar(&host, "host", "0.0.0.0", "Specify the server host")
	flag.StringVar(&port, "p", "80", "Specify the server port")
	flag.StringVar(&port, "port", "80", "Specify the server port")
	flag.Parse()

	// Use the flag contents to determine our handler method
	if serve {
		Serve(host, port)
	} else if len(file) > 0 {
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

// Serve launches a web server to process API requests
func Serve(host, port string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/freq", apiFreqHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: mux,
	}

	log.Printf("Listening on https://%v:%v", host, port)
	log.Fatal(server.ListenAndServe())
}

// Handle an API request
func apiFreqHandler(writer http.ResponseWriter, request *http.Request) {

	ls := NewLetterSet()
	CountLetters(request.FormValue("text"), &ls)
	ls.text = request.FormValue("text")

	writer.Header().Set("Content-Type", "application/json")
	jsonData, err := ls.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(writer, string(jsonData))
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
	file   string `json:"-"`
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

	// Display our letter counts in alphabeticall order
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

// MostCommonLetters returns a slice of strings containing
// the letters with the highest frequency count in the
// LetterSet, sorted alphabetically
func (l LetterSet) MostCommonLetters() []string {
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

// Empty indicates wether or not the letter set has any contents
func (l LetterSet) Empty() bool {
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
