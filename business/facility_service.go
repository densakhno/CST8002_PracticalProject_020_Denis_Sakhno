/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 15, 2025
Author: Denis Sakhno
Description: Business layer for NPRI facility management operations
This file contains the business logic for managing facility data in memory including CRUD operations.
*/

package business

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/denissakhno/CST8002_PracticalProject_020/models"
	"github.com/denissakhno/CST8002_PracticalProject_020/persistence"
)

// FacilityService provides business logic operations for facility management
// Manages the in-memory data structure and coordinates with persistence layer
type FacilityService struct {
	facilities []*models.Facility     // In-memory data structure (slice/array) storing facility objects
	repository persistence.FacilityRepo // FacilityRepo interface
}

// NewFacilityService creates a new facility service instance
// Parameters:
//   repository - repository for data persistence operations
// Returns:
//   *FacilityService - new service instance with empty facility collection
func NewFacilityService(repository persistence.FacilityRepo) *FacilityService {
	return &FacilityService{
		facilities: make([]*models.Facility, 0),
		repository: repository,
	}
}

// LoadData loads facility data from the data source using the persistence layer
// Returns:
//   error - any error encountered during data loading
// Replaces existing in-memory data with fresh data from file
func (fs *FacilityService) LoadData() error {
	facilities, err := fs.repository.LoadFacilities()
	if err != nil {
		return fmt.Errorf("failed to load facilities: %w", err)
	}

	fs.facilities = facilities
	return nil
}

// ReloadData reloads facility data from the data source, replacing in-memory data
// Returns:
//   int - number of facilities loaded
//   error - any error encountered during data reloading
// Provides functionality to refresh data from the original dataset
func (fs *FacilityService) ReloadData() (int, error) {
	if err := fs.LoadData(); err != nil {
		return 0, err
	}
	return len(fs.facilities), nil
}

// GetAllFacilities returns all facilities in the data structure
// Returns:
//   []*models.Facility - slice of all facility objects in memory
func (fs *FacilityService) GetAllFacilities() []*models.Facility {
	return fs.facilities
}

// GetFacilityCount returns the number of facilities in memory
// Returns:
//   int - count of facilities in the data structure
func (fs *FacilityService) GetFacilityCount() int {
	return len(fs.facilities)
}

// GetFacilityByIndex retrieves a facility by its index in the data structure
// Parameters:
//   index - zero-based index of the facility
// Returns:
//   *models.Facility - facility at the specified index
//   error - error if index is out of bounds
// Uses decision structure to validate array bounds
func (fs *FacilityService) GetFacilityByIndex(index int) (*models.Facility, error) {
	if index < 0 || index >= len(fs.facilities) {
		return nil, fmt.Errorf("index %d out of bounds (0-%d)", index, len(fs.facilities)-1)
	}
	return fs.facilities[index], nil
}

// GetFacilitiesByProvince returns facilities filtered by province
// Parameters:
//   province - province name to filter by
// Returns:
//   []*models.Facility - slice of facilities in the specified province
// Uses loop structure to iterate through facilities and decision structure for filtering
func (fs *FacilityService) GetFacilitiesByProvince(province string) []*models.Facility {
	var result []*models.Facility
	province = strings.ToLower(strings.TrimSpace(province))

	// Loop structure: iterate through all facilities
	for _, facility := range fs.facilities {
		// Decision structure: check if facility matches province
		if strings.ToLower(facility.Province) == province {
			result = append(result, facility)
		}
	}

	return result
}

// SearchFacilities searches for facilities by name (case-insensitive partial match)
// Parameters:
//   searchTerm - term to search for in facility names
// Returns:
//   []*models.Facility - slice of facilities matching the search term
// Uses loop and decision structures for searching
func (fs *FacilityService) SearchFacilities(searchTerm string) []*models.Facility {
	var result []*models.Facility
	searchTerm = strings.ToLower(strings.TrimSpace(searchTerm))

	if searchTerm == "" {
		return result
	}

	// Loop structure: iterate through all facilities
	for _, facility := range fs.facilities {
		// Decision structure: check multiple fields for matches
		if strings.Contains(strings.ToLower(facility.FacilityName), searchTerm) ||
		   strings.Contains(strings.ToLower(facility.CompanyName), searchTerm) ||
		   strings.Contains(strings.ToLower(facility.City), searchTerm) {
			result = append(result, facility)
		}
	}

	return result
}

// AddFacility adds a new facility to the in-memory data structure
// Parameters:
//   facility - facility object to add
// Returns:
//   error - validation error if facility is invalid
// Validates facility before adding to maintain data integrity
func (fs *FacilityService) AddFacility(facility *models.Facility) error {
	if facility == nil {
		return fmt.Errorf("facility cannot be nil")
	}

	// Validate facility data using model's validation method
	if err := facility.Validate(); err != nil {
		return fmt.Errorf("facility validation failed: %w", err)
	}

	// Check for duplicate NPRI ID using decision structure
	for _, existingFacility := range fs.facilities {
		if existingFacility.NPRIID == facility.NPRIID {
			return fmt.Errorf("facility with NPRI ID %s already exists", facility.NPRIID)
		}
	}

	// Add to data structure (array/slice)
	fs.facilities = append(fs.facilities, facility)
	return nil
}

// UpdateFacility updates an existing facility in memory and DB
// Parameters:
//   index - index of facility to update
//   fieldName - name of field to update
//   newValue - new value for the field
// Returns:
//   error - error if index is invalid or update fails
// Uses decision structure for validation and delegates field update to model
func (fs *FacilityService) UpdateFacility(index int, fieldName, newValue string) error {
	// Decision structure: validate index bounds
	if index < 0 || index >= len(fs.facilities) {
		return fmt.Errorf("index %d out of bounds (0-%d)", index, len(fs.facilities)-1)
	}

    // Update in-memory
    facility := fs.facilities[index]
    if err := facility.UpdateField(fieldName, newValue); err != nil {
        return err
    }
    // Write back to repo/db
    return fs.repository.UpdateAllFacilityInfo(facility)
}

// DeleteFacility removes a facility from memory and repository
// Parameters:
//   index - index of facility to delete
// Returns:
//   error - error if index is out of bounds
// Uses decision structure for validation and array manipulation
func (fs *FacilityService) DeleteFacility(index int) error {
	// Decision structure: validate index bounds
	if index < 0 || index >= len(fs.facilities) {
		return fmt.Errorf("index %d out of bounds (0-%d)", index, len(fs.facilities)-1)
	}

	npriid := fs.facilities[index].NPRIID
    if err := fs.repository.DeleteFacilityByID(npriid); err != nil {
        return err
    }

	// Remove facility from slice using array manipulation
	fs.facilities = append(fs.facilities[:index], fs.facilities[index+1:]...)
	return nil
}

// GetFacilitiesWithEmissions returns facilities that have emissions data
// Returns:
//   []*models.Facility - slice of facilities with valid emissions information
// Uses loop and decision structures for filtering
func (fs *FacilityService) GetFacilitiesWithEmissions() []*models.Facility {
	var result []*models.Facility

	// Loop structure: iterate through all facilities
	for _, facility := range fs.facilities {
		// Decision structure: check if facility has emissions data
		if facility.HasEmissionsData() {
			result = append(result, facility)
		}
	}

	return result
}

// GetTopEmitters returns facilities with highest emissions (up to specified limit)
// Parameters:
//   limit - maximum number of facilities to return
// Returns:
//   []*models.Facility - slice of top emitting facilities
// Uses loop structures for sorting and limiting results
func (fs *FacilityService) GetTopEmitters(limit int) []*models.Facility {
	// Get facilities with emissions data
	emissionFacilities := fs.GetFacilitiesWithEmissions()

	if len(emissionFacilities) == 0 {
		return []*models.Facility{}
	}

	// Simple bubble sort by emissions (loop structure with decision structure)
	for i := 0; i < len(emissionFacilities)-1; i++ {
		for j := 0; j < len(emissionFacilities)-i-1; j++ {
			emissions1, err1 := strconv.ParseFloat(emissionFacilities[j].Emissions, 64)
			emissions2, err2 := strconv.ParseFloat(emissionFacilities[j+1].Emissions, 64)

			// Decision structure: compare emissions values
			if err1 == nil && err2 == nil && emissions1 < emissions2 {
				// Swap facilities
				emissionFacilities[j], emissionFacilities[j+1] = emissionFacilities[j+1], emissionFacilities[j]
			}
		}
	}

	// Return limited results using decision structure
	if limit > len(emissionFacilities) {
		limit = len(emissionFacilities)
	}

	return emissionFacilities[:limit]
}

// GetDataSummary returns a summary of the current data in memory
// Returns:
//   string - formatted summary of facility data statistics
// Uses loop structures to calculate statistics
func (fs *FacilityService) GetDataSummary() string {
	totalFacilities := len(fs.facilities)
	if totalFacilities == 0 {
		return "No facilities loaded in memory."
	}

	// Count facilities by province using loop and decision structures
	provinceCount := make(map[string]int)
	facilitiesWithEmissions := 0

	for _, facility := range fs.facilities {
		provinceCount[facility.Province]++
		if facility.HasEmissionsData() {
			facilitiesWithEmissions++
		}
	}

	summary := fmt.Sprintf("Data Summary:\n")
	summary += fmt.Sprintf("- Total facilities: %d\n", totalFacilities)
	summary += fmt.Sprintf("- Facilities with emissions data: %d\n", facilitiesWithEmissions)
	summary += fmt.Sprintf("- Provinces represented: %d\n", len(provinceCount))

	return summary
}
