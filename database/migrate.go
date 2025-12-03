package database

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"sort"
)

func SortDirEntriesByName(entries []fs.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
}

//go:embed migrations/*.sql
var migrationsFiles embed.FS

func Migrate(db *sql.DB) error {
	entries, _ := migrationsFiles.ReadDir("migrations")

	SortDirEntriesByName(entries)

	for _, e := range entries {
		sqlContent, _ := migrationsFiles.ReadFile("migrations/" + e.Name())
		if _, err := db.Exec(string(sqlContent)); err != nil {
			log.Fatal("An error has occured: ", err)
			return err
		}
	}
	log.Println("Miration completed correctly")
	return nil
}
