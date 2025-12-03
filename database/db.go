package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/duckdb/duckdb-go/v2"
)

var dbIntance *sql.DB

func SetDB(db *sql.DB) {
	dbIntance = db
}

func DB() *sql.DB {
	return dbIntance
}

func InitDB(path string) error {
	db, err := sql.Open("duckdb", path)

	if err != nil {
		return fmt.Errorf("failed to open DuckDB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to DuckDB: %w", err)
	}

	SetDB(db)
	log.Println("DuckDB connected correctly")
	return nil
}

func CloseDB() error {
	if dbIntance != nil {
		return dbIntance.Close()
	}
	return nil
}
