/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 15, 2025
Author: Denis Sakhno
Description: Unit test for FacilityService.AddFacility
*/

package business

import (
	"testing"

	"github.com/denissakhno/CST8002_PracticalProject_020/models"
	"github.com/denissakhno/CST8002_PracticalProject_020/persistence"
)

// TestAddFacility ensures a valid facility is added to the service
func TestAddFacility(t *testing.T) {
	// Use in-memory "repository" that returns no data
	repo := persistence.NewFileRepository("")
	service := NewFacilityService(repo)

	// Create a minimal valid facility record
	mockFacitilityData := []string{"999", "Test Plant", "Test Co", "", "Test City", "TestProv", "A1A1A1", "0", "0", "100", "Tonnes (t)", "", "", "2022"}
	facility, err := models.NewFacility(mockFacitilityData)
	if err != nil {
		t.Fatalf("NewFacility failed: %v", err)
	}

	// Add to service
	if err := service.AddFacility(facility); err != nil {
		t.Errorf("AddFacility returned error: %v", err)
	}

	// Verify count
	facilityCount := service.GetFacilityCount()
	if facilityCount != 1 {
		t.Errorf("Expected 1 facility, got %d", service.GetFacilityCount())
	}
}
