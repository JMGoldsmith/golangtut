package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// expanding the struct
type ErrorRecord struct {
	Message string
	Count   int
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

	errorCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		errorPrefix := "[ERROR]"
		if strings.Contains(line, errorPrefix) {
			errorMessage := strings.SplitN(line, errorPrefix, 2)[1]
			errorCount[errorMessage]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var errors []ErrorRecord
	for message, count := range errorCount {
		errors = append(errors, ErrorRecord{Message: message, Count: count})
	}

	sort.Slice(errors, func(i, j int) bool {
		return errors[i].Count > errors[j].Count
	})

	for _, record := range errors {
		fmt.Printf("Encountered error %d times: %s\n", record.Count, record.Message)
	}
}
