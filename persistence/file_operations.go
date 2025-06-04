/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 15, 2025
Author: Denis Sakhno
Description: Persistence layer for NPRI facility data operations
This file handles all file I/O operations including reading CSV files and writing data with GUID filenames.
*/

package persistence

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/denissakhno/CST8002_PracticalProject_020/models"
	"github.com/google/uuid"
)

const PathToData = "data/data.csv"

// FileRepository provides file-based data persistence operations
// Handles reading from and writing to CSV files with proper error handling
type FileRepository struct {
	defaultFileName string // Default filename for reading data
}

// NewFileRepository creates a new file repository instance
// Parameters:
//   defaultFileName - default CSV file to read from
// Returns:
//   *FileRepository - new repository instance
func NewFileRepository(defaultFileName string) *FileRepository {
	return &FileRepository{
		defaultFileName: defaultFileName,
	}
}

// LoadFacilities reads facility data from the default CSV file
// Returns:
//   []*models.Facility - slice of facility pointers
//   error - any error encountered during file operations
// Uses exception handling to manage file access and parsing errors
func (fr *FileRepository) LoadFacilities() ([]*models.Facility, error) {
	return fr.LoadFacilitiesFromFile(fr.defaultFileName)
}

// LoadFacilitiesFromFile reads facility data from a specified CSV file
// Parameters:
//   filename - path to the CSV file containing NPRI facility data
// Returns:
//   []*models.Facility - slice of facility pointers
//   error - any error encountered during file operations or parsing
// Reads up to 100 records from the CSV file with proper error handling
func (fr *FileRepository) LoadFacilitiesFromFile(filename string) ([]*models.Facility, error) {
	// Open the CSV file with exception handling for missing files
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close() // Ensure file is closed when function exits

	// Create CSV reader using the encoding/csv API library
	reader := csv.NewReader(file)

	// Read all records from CSV file into memory
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %w", err)
	}

	// Validate that file contains data (at least header row)
	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	// Initialize slice to store facility objects
	var facilities []*models.Facility

	// Process up to 100 records (skip header at index 0)
	const maxRecords = 100
	recordCount := len(records) - 1 // Exclude header
	if recordCount > maxRecords {
		recordCount = maxRecords
	}

	// Process records with error handling for each facility
	for i := 1; i <= recordCount; i++ {
		facility, err := models.NewFacility(records[i])
		if err != nil {
			// Log error but continue processing other records
			fmt.Printf("Warning: Failed to create facility from record %d: %v\n", i, err)
			continue
		}
		facilities = append(facilities, facility)
	}

	if len(facilities) == 0 {
		return nil, fmt.Errorf("no valid facility records found in file")
	}

	return facilities, nil
}

// SaveFacilitiesToFile writes facility data to a CSV file with GUID filename
// Parameters:
//   facilities - slice of facility pointers to save
// Returns:
//   string - generated filename
//   error - any error encountered during file operations
// Uses UUID API to generate unique filenames for output files
func (fr *FileRepository) SaveFacilitiesToFile(facilities []*models.Facility) (string, error) {
	if len(facilities) == 0 {
		return "", fmt.Errorf("no facilities to save")
	}

	// Generate unique filename using UUID API
	guid := uuid.New()
	filename := fmt.Sprintf("npri_facilities_%s.csv", guid.String())

	// Create output file with proper error handling
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	// Create CSV writer using the encoding/csv API library
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{
		"NPRI ID", "Facility name", "Company name", "Address",
		"City", "Province", "PostalCode", "Latitude",
		"Longitude", "Emissions", "Units", "Facility details",
		"Facility information", "Report year",
	}

	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Write facility records using loop structure
	for _, facility := range facilities {
		record := facility.ToCSVRecord()
		if err := writer.Write(record); err != nil {
			return "", fmt.Errorf("failed to write facility record: %w", err)
		}
	}

	// Get absolute path for user feedback
	absPath, err := filepath.Abs(filename)
	if err != nil {
		absPath = filename // Fall back to relative path
	}

	return absPath, nil
}

// CheckFileExists validates if a file exists and is readable
// Parameters:
//   filename - path to file to check
// Returns:
//   bool - true if file exists and is readable, false otherwise
//   error - any error encountered during file stat operation
func (fr *FileRepository) CheckFileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("error checking file %s: %w", filename, err)
	}
	return true, nil
}

// GetFileInfo returns information about the default data file
// Returns:
//   string - file information including size and modification time
//   error - any error encountered during file stat operation
func (fr *FileRepository) GetFileInfo() (string, error) {
	info, err := os.Stat(fr.defaultFileName)
	if err != nil {
		return "", fmt.Errorf("failed to get file info for %s: %w", fr.defaultFileName, err)
	}

	return fmt.Sprintf("File: %s, Size: %d bytes, Modified: %s",
		fr.defaultFileName, info.Size(), info.ModTime().Format("2006-01-02 15:04:05")), nil
}
