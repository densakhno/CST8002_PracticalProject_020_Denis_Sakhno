/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: May 25, 2025
Author: Denis Sakhno
Description: Main program for reading NPRI facility pollution data
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// readFacilitiesFromCSV reads the CSV file and parses records into Facility structs
// Returns a slice of Facility objects and any error encountered
func readFacilitiesFromCSV(filename string) ([]Facility, error) {
	// Open the CSV file with exception handling
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %w", err)
	}

	// Check if file has data (at least header row)
	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	// Initialize slice to store facility objects
	var facilities []Facility

	// Parse first few records (skip header at index 0)
	maxRecords := 5 // Limit to first 5 data records
	for i := 1; i < len(records) && i <= maxRecords; i++ {
		record := records[i]

		// Ensure record has enough columns
		if len(record) >= 14 {
			facility := Facility{
				NPRIID:           record[0],
				FacilityName:     record[1],
				CompanyName:      record[2],
				Address:          record[3],
				City:             record[4],
				Province:         record[5],
				PostalCode:       record[6],
				Latitude:         record[7],
				Longitude:        record[8],
				Emissions:        record[9],
				Units:            record[10],
				FacilityDetails:  record[11],
				FacilityInfo:     record[12],
				ReportYear:       record[13],
			}
			facilities = append(facilities, facility)
		}
	}

	return facilities, nil
}

func main() {
	// Display program header
	fmt.Println("***NPRI Facility Data Reader***")
	fmt.Println("Loading facility data...")
	fmt.Println()

	// Read facilities from CSV with exception handling
	facilities, err := readFacilitiesFromCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Check the data file exists in the current directory.")
		return
	}

	fmt.Printf("Successfully loaded %d facility records.\n\n", len(facilities))
}