package database

import (
	"database/sql"
	"log"

	_ "github.com/duckdb/duckdb-go/v2"
)

var DB *sql.DB

func InitDB(path string) {
	var err error
	DB, err = sql.Open("duckdb", path)

	if err != nil {
		log.Fatalf("error opening DuckDB: %v", err)
	}

	log.Println("DuckDB connected correctly")
}
