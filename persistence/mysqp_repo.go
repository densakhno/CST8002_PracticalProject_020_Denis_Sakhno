package persistence

import (
	"database/sql"

	"github.com/denissakhno/CST8002_PracticalProject_020/models"
)

type DBrepository struct {
    db *sql.DB
}

const dbName = "cst8002"

func NewDBrepository(dsn string) (*DBrepository, error) {
    db, err := sql.Open(dbName, dsn)
    if err != nil {
        return nil, err
    }
    // Make sure the connection works
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return &DBrepository{db: db}, nil
}

func (mr *DBrepository) Close() error {
    if mr.db != nil {
        return mr.db.Close()
    }
    return nil
}

func (mr *DBrepository) LoadFacilities() ([]*models.Facility, error) {
    rows, err := mr.db.Query(
        `SELECT npriid, facilityname, companyname, address,
                city, province, postalcode, latitude, longitude,
                emissions, units, facilitydetails, facilityinfo, reportyear
         FROM facilities`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var facilities []*models.Facility
    for rows.Next() {
        var f models.Facility
        err := rows.Scan(
            &f.NPRIID, &f.FacilityName, &f.CompanyName, &f.Address,
            &f.City, &f.Province, &f.PostalCode, &f.Latitude, &f.Longitude,
            &f.Emissions, &f.Units, &f.FacilityDetails, &f.FacilityInfo, &f.ReportYear,
        )
        if err == nil {
            facilities = append(facilities, &f)
        }
    }
    return facilities, rows.Err()
}

func (mr *DBrepository) UpdateAllFacilityInfo(f *models.Facility) error {
    _, err := mr.db.Exec(`
        UPDATE facilities SET
            facilityname=?, companyname=?, address=?, city=?, province=?, postalcode=?,
            latitude=?, longitude=?, emissions=?, units=?,
            facilitydetails=?, facilityinfo=?, reportyear=?
        WHERE npriid=?`,
        f.FacilityName, f.CompanyName, f.Address, f.City, f.Province, f.PostalCode,
        f.Latitude, f.Longitude, f.Emissions, f.Units,
        f.FacilityDetails, f.FacilityInfo, f.ReportYear,
        f.NPRIID,
    )
    return err
}

func (mr *DBrepository) DeleteFacilityByID(id string) error {
    _, err := mr.db.Exec(
        `DELETE FROM facilities WHERE npriid = ?`, id,
    )
    return err
}

func (mr *DBrepository) AddFacility(f *models.Facility) error {
    _, err := mr.db.Exec(`
        INSERT INTO facilities
        (npriid, facilityname, companyname, address, city, province, postalcode,
         latitude, longitude, emissions, units, facilitydetails, facilityinfo, reportyear)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        f.NPRIID, f.FacilityName, f.CompanyName, f.Address, f.City, f.Province,
        f.PostalCode, f.Latitude, f.Longitude, f.Emissions, f.Units,
        f.FacilityDetails, f.FacilityInfo, f.ReportYear,
    )
    return err
}