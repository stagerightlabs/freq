package main

import (
	"flag"
	"fmt"
	"os"
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
		startServer(host, port)
	} else if len(file) > 0 {
		// i moved the error handling here, and made the func call return the error (if any)
		err := analyzeFile(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else if len(text) > 0 {
		analyzeText(text)
	} else {
		fmt.Println("You did not specify any input.")
		os.Exit(1) // this is an error state, right? so let's return a non zero value for Unix style error handling
	}
}
