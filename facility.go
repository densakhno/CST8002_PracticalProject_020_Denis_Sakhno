/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: May 25, 2025
Author: Denis Sakhno
Description: Facility record object for pollution data
*/

package main

import (
	"fmt"
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
