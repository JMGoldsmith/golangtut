package main

import (
	"fmt"
	"log"
	"os"
)

// the struct!
type ErrorRecord struct {
	Message string
}

// main function
func main() {
	// check if the log file exists
	if len(os.Args) != 2 {
		log.Fatal("Please provide a log file")
	}

	filename := os.Args[1]
	fmt.Printf(filename)
}
