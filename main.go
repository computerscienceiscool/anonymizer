// Description: This program reads a CSV file and anonymizes the specified columns using the specified techniques.
// Usage: go run main.go --input=input.csv --columns=column1,column2 --techniques=mask,pseudonymize

package main

import (
	"crypto/sha256"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Supported anonymization techniques
const (
	Mask         = "mask"
	Pseudonymize = "pseudonymize"
	Hash         = "hash"
	Generalize   = "generalize"
)

// Global pseudonym map
var pseudonymMap = make(map[string]string)

func main() {
	// Command-line flags
	inputFile := flag.String("input", "", "Path to the input CSV file")
	columns := flag.String("columns", "", "Comma-separated list of columns to anonymize")
	techniques := flag.String("techniques", "", "Comma-separated list of anonymization techniques (in order of columns)")
	flag.Parse()

	// Validate arguments
	if *inputFile == "" || *columns == "" || *techniques == "" {
		fmt.Println("Error: All arguments --input, --columns, and --techniques are required.")
		flag.Usage()
		return
	}

	columnList := strings.Split(*columns, ",")
	techniqueList := strings.Split(*techniques, ",")
	if len(columnList) != len(techniqueList) {
		fmt.Println("Error: Number of columns and techniques must match.")
		return
	}

	// Open the input file
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error: Unable to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse the CSV
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		fmt.Printf("Error: Unable to read headers: %v\n", err)
		return
	}

	// Map column names to indices
	columnIndices := map[string]int{}
	for i, header := range headers {
		columnIndices[header] = i
	}

	// Validate columns
	for _, column := range columnList {
		if _, exists := columnIndices[column]; !exists {
			fmt.Printf("Error: Column '%s' not found in CSV file.\n", column)
			return
		}
	}

	// Prepare output file
	outputFileName := strings.TrimSuffix(*inputFile, ".csv") + "_output.csv"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Error: Unable to create output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write headers to output
	if err := writer.Write(headers); err != nil {
		fmt.Printf("Error: Unable to write headers to output file: %v\n", err)
		return
	}

	// Anonymize rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: Unable to read record: %v\n", err)
			return
		}

		// Apply anonymization techniques
		for i, column := range columnList {
			index := columnIndices[column]
			technique := techniqueList[i]
			record[index] = anonymize(record[index], technique)
		}

		// Write anonymized record
		if err := writer.Write(record); err != nil {
			fmt.Printf("Error: Unable to write record to output file: %v\n", err)
			return
		}
	}

	fmt.Printf("Anonymized data written to %s\n", outputFileName)
}

// anonymize applies the specified technique to the data.
func anonymize(data, technique string) string {
	switch technique {
	case Mask:
		if len(data) > 2 {
			return data[:2] + strings.Repeat("*", len(data)-2)
		}
		return strings.Repeat("*", len(data))
	case Pseudonymize:
		if pseudonym, exists := pseudonymMap[data]; exists {
			return pseudonym
		}
		pseudonym := fmt.Sprintf("Person %d", rand.Intn(10000))
		pseudonymMap[data] = pseudonym
		return pseudonym
	case Hash:
		hash := sha256.Sum256([]byte(data))
		return fmt.Sprintf("%x", hash)
	case Generalize:
		if num, err := strconv.Atoi(data); err == nil {
			return fmt.Sprintf("%d-%d", (num/10)*10, (num/10)*10+9)
		}
		return "Unknown"
	default:
		return data
	}
}

// init seeds the random number generator.
func init() {
	rand.Seed(time.Now().UnixNano())
}
