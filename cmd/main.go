package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marceljsh/thrivepropcalc/internal/property"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	var records []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading input: %v", err)
		os.Exit(1)
	}

	if err := property.ProcessRecords(records); err != nil {
		fmt.Printf("error processing records: %v", err)
		os.Exit(1)
	}
}
