/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 15, 2025
Author: Denis Sakhno
Description: Presentation layer for NPRI facility management system
This file handles all user interactions and menu operations for the facility management application.
*/

package presentation

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/denissakhno/CST8002_PracticalProject_020/business"
	"github.com/denissakhno/CST8002_PracticalProject_020/models"
)

const studentName = "Denis Sakhno"
const studentID = "41121593"

// MenuSystem provides user interface operations for facility management
// Handles all user interactions and coordinates with business layer
type MenuSystem struct {
	service *business.FacilityService // Reference to business layer
	scanner *bufio.Scanner            // Scanner for user input
}

// NewMenuSystem creates a new menu system instance
// Parameters:
//   service - facility service from business layer
// Returns:
//   *MenuSystem - new menu system instance
func NewMenuSystem(service *business.FacilityService) *MenuSystem {
	return &MenuSystem{
		service: service,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// DisplayHeader shows the application header with author name
// Displays full name so it always remains visible as required
func (ms *MenuSystem) DisplayHeader() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              NPRI Facility Management System             ║")
	fmt.Printf("║              Author: %s                        ║\n", studentName)
		fmt.Printf("║              ID: %s                                ║\n", studentID)
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

// DisplayMainMenu shows the main menu options to the user
// Provides interactive options for all required functionality
func (ms *MenuSystem) DisplayMainMenu() {
	fmt.Println("═══ Main Menu ═══")
	fmt.Println("1. Reload data from dataset")
	fmt.Println("2. Display records")
	fmt.Println("3. Create new record")
	fmt.Println("4. Edit existing record")
	fmt.Println("5. Delete record")
	fmt.Println("6. Search facilities")
	fmt.Println("7. View data summary")
	fmt.Println("8. Exit")
	fmt.Printf("\nCurrent facilities in memory: %d\n", ms.service.GetFacilityCount())
	fmt.Print("Select an option (1-8): ")
}

// RunMainLoop executes the main application loop with user interactions
// Provides interactive functionality for all required operations
// Uses decision structure (switch) to handle menu selections
func (ms *MenuSystem) RunMainLoop() {
	// Display header with author name
	ms.DisplayHeader()

	fmt.Println("Loading initial data...")
	if err := ms.service.LoadData(); err != nil {
		fmt.Printf("Warning: Failed to load initial data: %v\n", err)
		fmt.Println("You can use 'Reload data' option to try loading again.")
	} else {
		fmt.Printf("Successfully loaded %d facilities.\n", ms.service.GetFacilityCount())
	}
	fmt.Println()

	// Main application loop structure
	for {
		ms.DisplayMainMenu()

		if !ms.scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(ms.scanner.Text())

		// Decision structure: handle user menu selections
		switch choice {
		case "1":
			ms.handleReloadData()
		case "2":
			ms.handleDisplayRecords()
		case "3":
			ms.handleCreateRecord()
		case "4":
			ms.handleEditRecord()
		case "5":
			ms.handleDeleteRecord()
		case "6":
			ms.handleSearchFacilities()
		case "7":
			ms.handleDataSummary()
		case "8":
			ms.handleExit()
			return
		default:
			fmt.Println("Invalid option. Please select 1-9.")
		}

		fmt.Println()
		ms.DisplayHeader() // Show name after each interaction
	}
}

// handleReloadData processes the reload data functionality
// Provides interactive option to reload data from dataset, replacing in-memory data
func (ms *MenuSystem) handleReloadData() {
	fmt.Println("\n=== Reload Data from Dataset ===")
	fmt.Print("This will replace all data in memory. Continue? (y/n): ")

	if !ms.scanner.Scan() {
		return
	}

	confirmation := strings.ToLower(strings.TrimSpace(ms.scanner.Text()))
	if confirmation != "y" && confirmation != "yes" {
		fmt.Println("Operation cancelled.")
		return
	}

	count, err := ms.service.ReloadData()
	if err != nil {
		fmt.Printf("Error reloading data: %v\n", err)
		return
	}

	fmt.Printf("Successfully reloaded %d facilities from dataset.\n", count)
}

// handleDisplayRecords provides options to display one or multiple records
// Allows user to select and display either single record or multiple records from in-memory data
func (ms *MenuSystem) handleDisplayRecords() {
	fmt.Println("\n=== Display Records ===")

	if ms.service.GetFacilityCount() == 0 {
		fmt.Println("No facilities loaded in memory.")
		return
	}

	fmt.Println("1. Display single record by index")
	fmt.Println("2. Display multiple records (all)")
	fmt.Println("3. Display multiple records (by province)")
	fmt.Println("4. Display top emitters")
	fmt.Print("Select display option (1-4): ")

	if !ms.scanner.Scan() {
		return
	}

	choice := strings.TrimSpace(ms.scanner.Text())

	// Decision structure for display options
	switch choice {
	case "1":
		ms.displaySingleRecord()
	case "2":
		ms.displayAllRecords()
	case "3":
		ms.displayRecordsByProvince()
	case "4":
		ms.displayTopEmitters()
	default:
		fmt.Println("Invalid display option.")
	}
}

// displaySingleRecord shows a single facility record by index
func (ms *MenuSystem) displaySingleRecord() {
	fmt.Printf("Enter record index (0-%d): ", ms.service.GetFacilityCount()-1)

	if !ms.scanner.Scan() {
		return
	}

	indexStr := strings.TrimSpace(ms.scanner.Text())
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Invalid index format.")
		return
	}

	facility, err := ms.service.GetFacilityByIndex(index)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\n--- Record %d ---\n", index)
	fmt.Print(facility.GetDetailedString())
}

// displayAllRecords shows all facility records using loop structure
func (ms *MenuSystem) displayAllRecords() {
	facilities := ms.service.GetAllFacilities()
	fmt.Printf("\n=== All Records (%d facilities) ===\n", len(facilities))

	// Loop structure: iterate through all facilities
	for i, facility := range facilities {
		fmt.Printf("%d. %s\n", i, facility.GetDisplayString())
	}
}

// displayRecordsByProvince shows facilities filtered by province
func (ms *MenuSystem) displayRecordsByProvince() {
	fmt.Print("Enter province name: ")

	if !ms.scanner.Scan() {
		return
	}

	province := strings.TrimSpace(ms.scanner.Text())
	facilities := ms.service.GetFacilitiesByProvince(province)

	if len(facilities) == 0 {
		fmt.Printf("No facilities found in province: %s\n", province)
		return
	}

	fmt.Printf("\n=== Facilities in %s (%d found) ===\n", province, len(facilities))

	// Loop structure: display filtered facilities
	for i, facility := range facilities {
		fmt.Printf("%d. %s\n", i+1, facility.GetDisplayString())
	}
}

// displayTopEmitters shows facilities with highest emissions
func (ms *MenuSystem) displayTopEmitters() {
	fmt.Print("Enter number of top emitters to display (default 10): ")

	if !ms.scanner.Scan() {
		return
	}

	limitStr := strings.TrimSpace(ms.scanner.Text())
	limit := 10 // Default value

	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	facilities := ms.service.GetTopEmitters(limit)

	if len(facilities) == 0 {
		fmt.Println("No facilities with emissions data found.")
		return
	}

	fmt.Printf("\n=== Top %d Emitters ===\n", len(facilities))

	// Loop structure: display top emitters
	for i, facility := range facilities {
		fmt.Printf("%d. %s\n", i+1, facility.GetDisplayString())
	}
}

// handleCreateRecord processes creating new records in memory
// Creates a new record and stores it in the data structure in memory
func (ms *MenuSystem) handleCreateRecord() {
	fmt.Println("\n=== Create New Record ===")
	fmt.Println("Enter facility information (press Enter for empty fields):")

	// Collect facility data from user
	data := make([]string, 14)
	fields := []string{
		"NPRI ID", "Facility name", "Company name", "Address",
		"City", "Province", "Postal code", "Latitude",
		"Longitude", "Emissions", "Units", "Facility details",
		"Facility information", "Report year",
	}

	// Loop structure: collect input for each field
	for i, field := range fields {
		fmt.Printf("%s: ", field)
		if ms.scanner.Scan() {
			data[i] = strings.TrimSpace(ms.scanner.Text())
		}
	}

	// Create facility using model layer
	facility, err := models.NewFacility(data)
	if err != nil {
		fmt.Printf("Error creating facility: %v\n", err)
		return
	}

	// Add to business layer data structure
	if err := ms.service.AddFacility(facility); err != nil {
		fmt.Printf("Error adding facility: %v\n", err)
		return
	}

	fmt.Println("✅ Facility successfully created and added to memory.")
	fmt.Printf("Total facilities in memory: %d\n", ms.service.GetFacilityCount())
}

// handleEditRecord processes editing existing records in memory
// Selects and edits a record held in the data structure in memory
func (ms *MenuSystem) handleEditRecord() {
	fmt.Println("\n=== Edit Existing Record ===")

	if ms.service.GetFacilityCount() == 0 {
		fmt.Println("No facilities available to edit.")
		return
	}

	// Select facility to edit
	fmt.Printf("Enter record index to edit (0-%d): ", ms.service.GetFacilityCount()-1)

	if !ms.scanner.Scan() {
		return
	}

	indexStr := strings.TrimSpace(ms.scanner.Text())
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Invalid index format.")
		return
	}

	facility, err := ms.service.GetFacilityByIndex(index)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Display current facility data
	fmt.Printf("\nCurrent facility data:\n%s\n", facility.GetDetailedString())

	// Show available fields for editing
	fmt.Println("Available fields to edit:")
	fieldNames := models.GetFieldNames()

	// Loop structure: display field options
	for i, fieldName := range fieldNames {
		fmt.Printf("%d. %s\n", i+1, fieldName)
	}

	fmt.Printf("Select field to edit (1-%d): ", len(fieldNames))

	if !ms.scanner.Scan() {
		return
	}

	fieldIndexStr := strings.TrimSpace(ms.scanner.Text())
	fieldIndex, err := strconv.Atoi(fieldIndexStr)
	if err != nil || fieldIndex < 1 || fieldIndex > len(fieldNames) {
		fmt.Println("Invalid field selection.")
		return
	}

	selectedField := fieldNames[fieldIndex-1]
	fmt.Printf("Enter new value for %s: ", selectedField)

	if !ms.scanner.Scan() {
		return
	}

	newValue := strings.TrimSpace(ms.scanner.Text())

	// Update facility using business layer
	if err := ms.service.UpdateFacility(index, selectedField, newValue); err != nil {
		fmt.Printf("Error updating facility: %v\n", err)
		return
	}

	fmt.Println("✅ Facility successfully updated in memory.")
}

// handleDeleteRecord processes deleting records from memory
// Selects and deletes a record from the data structure in memory
func (ms *MenuSystem) handleDeleteRecord() {
	fmt.Println("\n=== Delete Record ===")

	if ms.service.GetFacilityCount() == 0 {
		fmt.Println("No facilities available to delete.")
		return
	}

	// Show current facilities
	ms.displayAllRecords()

	fmt.Printf("\nEnter record index to delete (0-%d): ", ms.service.GetFacilityCount()-1)

	if !ms.scanner.Scan() {
		return
	}

	indexStr := strings.TrimSpace(ms.scanner.Text())
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Invalid index format.")
		return
	}

	facility, err := ms.service.GetFacilityByIndex(index)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Confirm deletion
	fmt.Printf("\nAre you sure you want to delete this facility?\n")
	fmt.Printf("NPRI ID: %s, Facility: %s\n", facility.NPRIID, facility.FacilityName)
	fmt.Print("Confirm deletion (y/n): ")

	if !ms.scanner.Scan() {
		return
	}

	confirmation := strings.ToLower(strings.TrimSpace(ms.scanner.Text()))
	if confirmation != "y" && confirmation != "yes" {
		fmt.Println("Deletion cancelled.")
		return
	}

	// Delete from business layer data structure
	if err := ms.service.DeleteFacility(index); err != nil {
		fmt.Printf("Error deleting facility: %v\n", err)
		return
	}

	fmt.Println("✅ Facility successfully deleted from memory.")
	fmt.Printf("Remaining facilities in memory: %d\n", ms.service.GetFacilityCount())
}

// handleSearchFacilities processes facility search functionality
func (ms *MenuSystem) handleSearchFacilities() {
	fmt.Println("\n=== Search Facilities ===")
	fmt.Print("Enter search term (facility name, company, or city): ")

	if !ms.scanner.Scan() {
		return
	}

	searchTerm := strings.TrimSpace(ms.scanner.Text())
	if searchTerm == "" {
		fmt.Println("Search term cannot be empty.")
		return
	}

	facilities := ms.service.SearchFacilities(searchTerm)

	if len(facilities) == 0 {
		fmt.Printf("No facilities found matching: %s\n", searchTerm)
		return
	}

	fmt.Printf("\n=== Search Results for '%s' (%d found) ===\n", searchTerm, len(facilities))

	// Loop structure: display search results
	for i, facility := range facilities {
		fmt.Printf("%d. %s\n", i+1, facility.GetDisplayString())
	}
}

// handleDataSummary displays summary statistics of current data
func (ms *MenuSystem) handleDataSummary() {
	fmt.Println("\n=== Data Summary ===")
	summary := ms.service.GetDataSummary()
	fmt.Println(summary)
}

// handleExit processes application exit
func (ms *MenuSystem) handleExit() {
	fmt.Println("\n=== Exit Application ===")
	fmt.Println("Thank you for using NPRI Facility Management System!")
	fmt.Println("Author: Denis Sakhno")
	fmt.Println("Program terminated successfully.")
}
