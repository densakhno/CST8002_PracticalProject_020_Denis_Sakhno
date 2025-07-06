/*
Course: CST 8002 Programming Language Research Project
Professor: Stanley Pieda, Tyler DeLay
Due Date: July 13, 2025
Author: Denis Sakhno
Description: Persistence layer for NPRI facility data operations
This file provides repository implementations for facility data storage.
*/

package persistence

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/denissakhno/CST8002_PracticalProject_020/models"
)

// DBrepository implements data persistence for facilities using a MySQL database.
type DBrepository struct {
    db *sql.DB
}

// NewDBrepository creates a new DBrepository instance and a connection to MySQL.
//
// Parameters:
//   - dsn: the MySQL Data Source Name (format: user:pass@tcp(host:port)/dbname)
//
// Returns:
//   - pointer to DBrepository
//   - error if connection could not be established
func NewDBrepository(dsn string) (*DBrepository, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    // Make sure the connection works
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return &DBrepository{db: db}, nil
}

// Close terminates the database connection for this repository.
//
// Returns:
//   - error if closing the DB fails
func (mr *DBrepository) Close() error {
    if mr.db != nil {
        return mr.db.Close()
    }
    return nil
}

// LoadFacilities retrieves all facility records from the MySQL database.
//
// Returns:
//   - slice of pointers to Facility
//   - error if SQL querying or scanning fails
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

// UpdateAllFacilityInfo updates all the fields of a facility in the database,
// matched by primary key (NPRI ID).
//
// Parameters:
//   - f: pointer to the Facility struct to update
//
// Returns:
//   - error if the update fails
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

// DeleteFacilityByID deletes a facility from the database using its NPRI ID.
//
// Parameters:
//   - id: the NPRI ID of the facility to delete
//
// Returns:
//   - error if the delete operation fails
func (mr *DBrepository) DeleteFacilityByID(id string) error {
    _, err := mr.db.Exec(
        `DELETE FROM facilities WHERE npriid = ?`, id,
    )
    return err
}

// AddFacility inserts a new facility into the database.
//
// Parameters:
//   - f: pointer to the new Facility struct to insert
//
// Returns:
//   - error if the insert operation fails
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