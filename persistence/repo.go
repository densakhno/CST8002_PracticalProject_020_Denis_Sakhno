/*
Course: CST 8002 Programming Language Research Project
Professor: Stanley Pieda, Tyler DeLay
Due Date: July 13, 2025
Author: Denis Sakhno
Description: Persistence layer for NPRI facility data operations
This file defines interfaces and implementations for data persistence.
*/

package persistence

import "github.com/denissakhno/CST8002_PracticalProject_020/models"

// FacilityRepo defines the contract for persistent facility storage backends.
type FacilityRepo interface {
    LoadFacilities() ([]*models.Facility, error)
    AddFacility(f *models.Facility) error
    UpdateAllFacilityInfo(f *models.Facility) error
    DeleteFacilityByID(npriid string) error
}
