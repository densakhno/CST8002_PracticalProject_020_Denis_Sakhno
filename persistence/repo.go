package persistence

import "github.com/denissakhno/CST8002_PracticalProject_020/models"

type FacilityRepo interface {
    LoadFacilities() ([]*models.Facility, error)
    AddFacility(f *models.Facility) error
    UpdateAllFacilityInfo(f *models.Facility) error
    DeleteFacilityByID(npriid string) error
}
