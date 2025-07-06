# NPRI Facility Management System

**Course:** CST 8002 Programming Language Research Project
**Author:** Denis Sakhno
**Due Date:** July 13, 2025

## Overview

This project is a command-line application for managing NPRI (National Pollutant Release Inventory) facility pollution data. It supports loading, viewing, searching, editing, and persisting facility records using both CSV files and a MySQL database.

## Features

- Load facility data from CSV or MySQL
- Display, search, and filter facilities (by province, top emitters, etc.)
- Add, edit, and delete facility records
- View data summaries (counts, emissions, provinces)
- Import CSV data into MySQL

## Data Model

Each facility record includes:

- NPRI ID
- Facility name
- Company name
- Address
- City
- Province
- Postal code
- Latitude, Longitude
- Emissions, Units
- Facility details, Facility information
- Report year

## Setup

1. **Go Version:** 1.23.1 or later
2. **Dependencies:**
   - github.com/go-sql-driver/mysql
3. **Database:**
   - MySQL server required
   - Use `sql_scripts.txt` to create the database and `facilities` table:
     ```sql
     CREATE DATABASE IF NOT EXISTS cst8002 DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
     USE cst8002;
     CREATE TABLE IF NOT EXISTS facilities (
         npriid VARCHAR(64) PRIMARY KEY,
         facilityname VARCHAR(255),
         companyname VARCHAR(255),
         address VARCHAR(255),
         city VARCHAR(100),
         province VARCHAR(100),
         postalcode VARCHAR(20),
         latitude VARCHAR(50),
         longitude VARCHAR(50),
         emissions VARCHAR(50),
         units VARCHAR(50),
         facilitydetails TEXT,
         facilityinfo TEXT,
         reportyear VARCHAR(10)
     );
     ```
4. **CSV Data:** Place your data file at `data/data.csv` with the following header:
   ```csv
   NPRI ID,Facility name,Company name,Address,City,Province,PostalCode,Latitude,Longitude,Emissions,Units,Facility details,Facility information,Report year
   ```

## Usage

- Set the MySQL connection string in the environment variable `DB_DSN` (e.g., `export DB_DSN="user:password@tcp(localhost:3306)/cst8002"`).
- Run the program:
  ```sh
  go run main.go
  ```
- To import CSV data into MySQL and exit:
  ```sh
  go run main.go -import-csv
  ```
- Follow the interactive menu to manage facilities.

## Project Structure

- `main.go` — Entry point, CLI and import logic
- `business/` — Business logic (CRUD, search, summary)
- `models/` — Facility data model
- `persistence/` — File and MySQL data access
- `presentation/` — CLI menu and user interaction
- `data/` — CSV data files

## License

For educational use only.
