/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: July 13, 2025
Author: Denis Sakhno
Description: Unit test for adding a new facility to DB
*/

package business

import (
	"os"
	"testing"

	"github.com/denissakhno/CST8002_PracticalProject_020/models"
	"github.com/denissakhno/CST8002_PracticalProject_020/persistence"
)

// Test adding a facility to DB
func TestAddFacilityMySQL(t *testing.T) {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        t.Skip("DB_DSN not set; skipping test")
    }
    repo, err := persistence.NewDBrepository(dsn)
    if err != nil {
        t.Fatalf("Failed to connect to MySQL: %v", err)
    }
    defer repo.Close()

    service := NewFacilityService(repo)

    // Clean up for future test rerun
    _ = repo.DeleteFacilityByID("888")

    mockFacilityData := []string{"888", "Test Facility", "Test Comp", "", "Test City", "TestProv", "K1S3h8", "0", "0", "100", "Tonnes (t)", "", "", "2025"}
    facility, err := models.NewFacility(mockFacilityData)
    if err != nil {
        t.Fatalf("NewFacility failed: %v", err)
    }

    if err := service.AddFacility(facility); err != nil {
        t.Errorf("AddFacility returned error: %v", err)
    }

    // Reload the data from DB
    _ = service.LoadData()
    facilityCount := 0
    for _, f := range service.GetAllFacilities() {
        if f.NPRIID == "888" {
            facilityCount++
        }
    }
    if facilityCount != 1 {
        t.Errorf("Expected 1 facility with NPRIID=888, got %d", facilityCount)
    }

    // Clean up
    _ = repo.DeleteFacilityByID("888")
}