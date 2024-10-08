package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// TODO you have to make a new list

// Function to convert JSON to CSV
func jsonToCSV(jsonfile, csvFile string) error {
	data, err := Reader(jsonfile)
	if err != nil {
		return fmt.Errorf("fialed to unmarsh JSON file: %v", err)
	}

	// UNmarshal Json in to slice of maps
	var jsonData []map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Create CSV file
	csvOutput, err := os.Create(csvFile)
	if err != nil {
		return fmt.Errorf("fialed to create CSV file: %v", err)
	}
	defer csvOutput.Close()

	writer := csv.NewWriter(csvOutput)
	defer writer.Flush()

	// Write CSV header
	var header []string
	for key := range jsonData[0] {
		header = append(header, header, key)
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("fialed to write CSV header: %v", err)
	}

	// write CSV data
	for _, recored := range jsonData {
		row := make([]string, len(header))
		for i, header := range headers {
			row[i] = fmt.Sprintf("%v", recored[header])
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("fialed to write writer CSV recored: %v", err)
		}
	}
	return nil
}

// For converting CSV to Json
func csvToJSON(csvFile, jsonFiel string) error {
	// Read CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Read all records form CSV
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("fialed to read CSV recoreds: %v", err)
	}

	// Prepare Json data
	var jsonData []map[string]interface{}
	headers := records[0]

	for _, record := range records[1:] {
		entry := make(map[string]interface{})
		for i, value := range record {
			entry[headers[i]] = value
		}
		jsonData = append(jsonData, entry)
	}
	// Write JSON to file
	if err := io.WriteFile(jsonFile, JsonOutput, 0644); err != nil {
		return fmt.Errorf("failed to write JSON to file: %v", err)
	}
	return nil
}
