package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func openFile(filename string) (file *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Couldn't open the file you provided: \"%s\"\n", filename))
	}
	return
}

func parseCSV(file *os.File) (csvContent [][]string) {
	csvFile := csv.NewReader(file)
	csvContent, err := csvFile.ReadAll()
	if err != nil {
		exit("Couldn't parse the CSV file")
	}
	return
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
