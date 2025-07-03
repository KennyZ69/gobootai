package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	verbose = flag.Bool("v", false, "Enable verbose output")
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	prompt := os.Args[1]
	flag.Parse()

	resp, err := GenerateResponse(prompt, *verbose)
	if err != nil {
		fmt.println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Response:\n%v\n", resp)
}

func usage() {
	println("Usage: gobootai <prompt> [--verbose]")
	os.Exit(1)
}
