package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

/*
type DataSource struct {
	Name         string `json:"name"`
	AuthStrategy string `json:"authstrategy"`
	Location     string `json:"location"`
	FriendlyName string `json:"friendly_name"`
}

type Credentials struct {
	AuthStrategy string `json:"authstrategy"`
}
*/

// TODO: make a cvs funictions to open and edit CSV file
func readCSV(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// The CSV reader
func parseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

// Iterate through the
func processCSV(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		fmt.Println(record)
	}
}

// Modify data to remove errors
func modifyCSV()

// TODO: Pull data from CSV file and add the information to a struct.
// TODO: next step is to fix the object error

func main() {
	data, err := readCSV("testdata.csv")
	if err != nil {
		fmt.Println("Error readign file:", err)
		return
	}
	reader, err := parseCSV(data)
	if err != nil {
		fmt.Println("Error creating csv reader:", err)
		return
	}
	processCSV(reader)
}
