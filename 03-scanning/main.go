package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ErrorRecord struct {
	Message string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a log file")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// adding a scanner to scan the file and add errors.
	scanner := bufio.NewScanner(file)
	var errorMessage []string
	for scanner.Scan() {
		line := scanner.Text()
		errorPrefix := "[ERROR]"
		if strings.Contains(line, errorPrefix) {
			errorMessage = append(errorMessage, strings.SplitN(line, errorPrefix, 2)[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range errorMessage {
		fmt.Printf("%s\n", v)
	}

}
