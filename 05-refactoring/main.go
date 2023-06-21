package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type ErrorRecord struct {
	Message    string
	Count      int
	LineNumber int
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
	scanned := scanFile(file)
	errorList := countErrors(scanned)
	sortAndPrint(errorList)

}

func scanFile(file *os.File) map[string][]int {
	errorCount := make(map[string][]int)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		errorPrefix := "[ERROR]"
		if strings.Contains(line, errorPrefix) {
			errorMessage := strings.SplitN(line, errorPrefix, 2)[1]
			if nil == errorCount[errorMessage] {
				errorCount[errorMessage] = []int{0, 0}
			}
			currentCounter := errorCount[errorMessage][0]
			currentCounter++
			errorCount[errorMessage] = []int{currentCounter, lineNumber}
			// errorCount[errorMessage]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return errorCount
}

func countErrors(scanned map[string][]int) []ErrorRecord {
	var errors []ErrorRecord
	for message, count := range scanned {
		errors = append(errors, ErrorRecord{Message: message, Count: count[0], LineNumber: count[1]})
	}
	return errors
}

func sortAndPrint(errorList []ErrorRecord) {
	sort.Slice(errorList, func(i, j int) bool {
		return errorList[i].Count > errorList[j].Count
	})

	for _, record := range errorList {
		fmt.Printf("On Line %d Encountered error %d times: %s\n", record.LineNumber, record.Count, record.Message)
	}
}
