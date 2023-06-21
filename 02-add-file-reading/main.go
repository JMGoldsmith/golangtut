package main

import (
	"fmt"
	"log"
	"os"
)

type ErrorRecord struct {
	Message string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a log file")
	}

	filename := os.Args[1]
	// opening the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// this is important!
	defer file.Close()
	// add print here for file
	fmt.Printf(filename)
}
