/*
Course: CST 8002 Programming Language Research Project
Professor: Stanley Pieda, Tyler DeLay
Due Date: July 13, 2025
Author: Denis Sakhno
Description: Facility model/entity representing NPRI pollution facility records
This file contains the data structure definition for facility records with validation methods
*/

package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Facility represents a pollution facility record from the NPRI dataset
// This struct uses the column names from the CSV as field names
type Facility struct {
	NPRIID           string
	FacilityName     string
	CompanyName      string
	Address          string
	City             string
	Province         string
	PostalCode       string
	Latitude         string
	Longitude        string
	Emissions        string
	Units            string
	FacilityDetails  string
	FacilityInfo     string
	ReportYear       string
}

// GetDisplayString returns a formatted string representation of the facility
func (f *Facility) GetDisplayString() string {
	return fmt.Sprintf("ID: %s | %s | %s, %s | Emissions: %s %s",
		f.NPRIID, f.FacilityName, f.City, f.Province, f.Emissions, f.Units)
}

// GetDetailedString returns a comprehensive string representation of the facility
// Returns all facility information formatted for detailed display
func (f *Facility) GetDetailedString() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("NPRI ID: %s\n", f.NPRIID))
	builder.WriteString(fmt.Sprintf("Facility: %s\n", f.FacilityName))
	builder.WriteString(fmt.Sprintf("Company: %s\n", f.CompanyName))

	if f.Address != "" {
		builder.WriteString(fmt.Sprintf("Address: %s\n", f.Address))
	}
	builder.WriteString(fmt.Sprintf("Location: %s, %s %s\n", f.City, f.Province, f.PostalCode))

	if f.Latitude != "" && f.Longitude != "" {
		builder.WriteString(fmt.Sprintf("Coordinates: %s, %s\n", f.Latitude, f.Longitude))
	}

	if f.Emissions != "" {
		builder.WriteString(fmt.Sprintf("Emissions: %s %s\n", f.Emissions, f.Units))
	}

	builder.WriteString(fmt.Sprintf("Report Year: %s\n", f.ReportYear))

	return builder.String()
}

// NewFacility creates a new Facility instance with validation
// Parameters:
//   data - slice of strings containing facility data from CSV
// Returns:
//   *Facility - pointer to new facility instance
//   error - validation error if data is invalid
func NewFacility(data []string) (*Facility, error) {
	if len(data) < 14 {
		return nil, fmt.Errorf("insufficient data fields: expected 14, got %d", len(data))
	}

	facility := &Facility{
		NPRIID:           strings.TrimSpace(data[0]),
		FacilityName:     strings.TrimSpace(data[1]),
		CompanyName:      strings.TrimSpace(data[2]),
		Address:          strings.TrimSpace(data[3]),
		City:             strings.TrimSpace(data[4]),
		Province:         strings.TrimSpace(data[5]),
		PostalCode:       strings.TrimSpace(data[6]),
		Latitude:         strings.TrimSpace(data[7]),
		Longitude:        strings.TrimSpace(data[8]),
		Emissions:        strings.TrimSpace(data[9]),
		Units:            strings.TrimSpace(data[10]),
		FacilityDetails:  strings.TrimSpace(data[11]),
		FacilityInfo:     strings.TrimSpace(data[12]),
		ReportYear:       strings.TrimSpace(data[13]),
	}

	return facility, facility.Validate()
}

// ToCSVRecord converts the facility to a CSV record format
// Returns a slice of strings representing the facility data for CSV output
func (f *Facility) ToCSVRecord() []string {
	return []string{
		f.NPRIID,
		f.FacilityName,
		f.CompanyName,
		f.Address,
		f.City,
		f.Province,
		f.PostalCode,
		f.Latitude,
		f.Longitude,
		f.Emissions,
		f.Units,
		f.FacilityDetails,
		f.FacilityInfo,
		f.ReportYear,
	}
}

// Validate performs validation on facility data
// Returns error if required fields are missing or invalid
func (f *Facility) Validate() error {
	if f.NPRIID == "" {
		return fmt.Errorf("NPRI ID is required")
	}
	if f.FacilityName == "" {
		return fmt.Errorf("facility name is required")
	}
	if f.City == "" {
		return fmt.Errorf("city is required")
	}
	if f.Province == "" {
		return fmt.Errorf("province is required")
	}

	// Validate emissions if present
	if f.Emissions != "" {
		if _, err := strconv.ParseFloat(f.Emissions, 64); err != nil {
			return fmt.Errorf("invalid emissions value: %s", f.Emissions)
		}
	}

	return nil
}

// HasEmissionsData checks if the facility has valid emissions information
// Returns true if both emissions value and units are present, false otherwise
// Useful for filtering facilities with complete pollution data
func (f *Facility) HasEmissionsData() bool {
	return f.Emissions != "" && f.Units != ""
}

// UpdateField updates a specific field of the facility
// Parameters:
//   fieldName - name of the field to update
//   value - new value for the field
// Returns error if field name is invalid
func (f *Facility) UpdateField(fieldName, value string) error {
	value = strings.TrimSpace(value)

	switch strings.ToLower(fieldName) {
	case "npriid":
		f.NPRIID = value
	case "facilityname":
		f.FacilityName = value
	case "companyname":
		f.CompanyName = value
	case "address":
		f.Address = value
	case "city":
		f.City = value
	case "province":
		f.Province = value
	case "postalcode":
		f.PostalCode = value
	case "latitude":
		f.Latitude = value
	case "longitude":
		f.Longitude = value
	case "emissions":
		f.Emissions = value
	case "units":
		f.Units = value
	case "facilitydetails":
		f.FacilityDetails = value
	case "facilityinfo":
		f.FacilityInfo = value
	case "reportyear":
		f.ReportYear = value
	default:
		return fmt.Errorf("invalid field name: %s", fieldName)
	}

	return f.Validate()
}

// GetFieldNames returns a list of all available field names for editing
// Returns slice of field names that can be used with UpdateField method
func GetFieldNames() []string {
	return []string{
		"npriid", "facilityname", "companyname", "address",
		"city", "province", "postalcode", "latitude",
		"longitude", "emissions", "units", "facilitydetails",
		"facilityinfo", "reportyear",
	}
}