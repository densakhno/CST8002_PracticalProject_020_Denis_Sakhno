/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 15, 2025
Author: Denis Sakhno
Description: Main program for reading NPRI facility pollution data
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/denissakhno/CST8002_PracticalProject_020/business"
	"github.com/denissakhno/CST8002_PracticalProject_020/persistence"
	"github.com/denissakhno/CST8002_PracticalProject_020/presentation"
)


func importCSVtoMySQL(csvRepo *persistence.FileRepository, mysqlRepo *persistence.DBrepository) error {
    facilities, err := csvRepo.LoadFacilities()
    if err != nil {
        return err
    }
    for _, f := range facilities {
        if err := mysqlRepo.AddFacility(f); err != nil {
            fmt.Printf("Failed to import NPRIID %s: %v\n", f.NPRIID, err)
        }
    }
    fmt.Printf("Imported %d facilities to MySQL.\n", len(facilities))
    return nil
}


func main() {
    importCSV := flag.Bool("import-csv", false, "Import CSV to MySQL, then exit.")
    flag.Parse()

    // For import CSV migration
    if *importCSV {
        fileRepo := persistence.NewFileRepository(persistence.PathToData)
        dsn := os.Getenv("DB_DSN")
        if dsn == "" {
            fmt.Println("The DB_DSN env variable is not set to connect to MySQL.")
            os.Exit(1)
        }
        mysqlRepo, err := persistence.NewDBrepository(dsn)
        if err != nil {
            fmt.Printf("Error connecting to MySQL: %v\n", err)
            os.Exit(1)
        }
        defer mysqlRepo.Close()

        if err := importCSVtoMySQL(fileRepo, mysqlRepo); err != nil {
            fmt.Printf("CSV import failed: %v\n", err)
            os.Exit(1)
        }
        os.Exit(0)
    }

    // Normal application mode
	var repo persistence.FacilityRepo
    var err error

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		fmt.Println("Set the DB_DSN env variable to connect to MySQL.")
		os.Exit(1)
	}
	repo, err = persistence.NewDBrepository(dsn)
	if err != nil {
		fmt.Printf("Error connecting to MySQL: %v\n", err)
		os.Exit(1)
	}
	defer repo.(*persistence.DBrepository).Close()

	// Initialize business layer
	service := business.NewFacilityService(repo)

	// Initialize presentation layer
	menu := presentation.NewMenuSystem(service)
	menu.RunMainLoop()
}