package persistence

import (
	"database/sql"
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
