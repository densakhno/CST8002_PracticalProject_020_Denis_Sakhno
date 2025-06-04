/*
Course: CST 8333 Programming Language Research Project
Professor: Stanley Pieda
Due Date: May 25, 2025
Author: Denis Sakhno
Description: Main program for reading NPRI facility pollution data
*/

package main

import (
	"fmt"
	"os"

	"github.com/denissakhno/CST8002_PracticalProject_020/business"
	"github.com/denissakhno/CST8002_PracticalProject_020/persistence"
	"github.com/denissakhno/CST8002_PracticalProject_020/presentation"
)

func main() {
	// Ensure data file exists
	repo := persistence.NewFileRepository("data/data.csv")
	ok, err := repo.CheckFileExists("data/data.csv")
	if err != nil {
		fmt.Printf("Error checking data file: %v\n", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Println("Dataset file not found. Please place it in the data/ folder.")
		os.Exit(1)
	}

	// Initialize business layer
	service := business.NewFacilityService(repo)

	// Initialize presentation layer
	menu := presentation.NewMenuSystem(service)
	menu.RunMainLoop()
}