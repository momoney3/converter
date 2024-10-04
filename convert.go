package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Function to convert JSON to CSV
func jsonToCSV(jsonFile string, csvFile string) error {
	// Read JSON file
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	// Unmarshal JSON into a slice of maps
	var jsonData []map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Create CSV file
	csvOutput, err := os.Create(csvFile)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer csvOutput.Close()

	writer := csv.NewWriter(csvOutput)
	defer writer.Flush()

	// Write CSV header
	var headers []string
	for key := range jsonData[0] {
		headers = append(headers, key)
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV header: %v", err)
	}

	// Write CSV data
	for _, record := range jsonData {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = fmt.Sprintf("%v", record[header])
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV record: %v", err)
		}
	}

	return nil
}

// Function to convert CSV to JSON
func csvToJSON(csvFile string, jsonFile string) error {
	// Read CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV records: %v", err)
	}

	// Prepare JSON data
	var jsonData []map[string]interface{}
	headers := records[0]

	for _, record := range records[1:] {
		entry := make(map[string]interface{})
		for i, value := range record {
			entry[headers[i]] = value
		}
		jsonData = append(jsonData, entry)
	}

	// Convert to JSON
	jsonOutput, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Write JSON to file
	if err := ioutil.WriteFile(jsonFile, jsonOutput, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run convert.go <input file> <output file> <format: json-to-csv|csv-to-json>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	format := strings.ToLower(os.Args[3])

	var err error
	switch format {
	case "json-to-csv":
		err = jsonToCSV(inputFile, outputFile)
	case "csv-to-json":
		err = csvToJSON(inputFile, outputFile)
	default:
		fmt.Println("Invalid format. Use 'json-to-csv' or 'csv-to-json'.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Conversion successful!")
	}
}
