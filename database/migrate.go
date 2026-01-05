package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"
)

func SortDirEntriesByName(entries []fs.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
}

//go:embed migrations/*.sql
var migrationsFiles embed.FS

func readMigrationFiles() ([]fs.DirEntry, error) {
	entries, err := migrationsFiles.ReadDir("migrations")
	if err != nil {
		return nil, fmt.Errorf("cannot read migrations directory: %w", err)
	}

	SortDirEntriesByName(entries)
	return entries, nil
}

func readSQLFile(name string) (string, error) {
	content, err := migrationsFiles.ReadFile("migrations/" + name)
	if err != nil {
		return "", fmt.Errorf("cannot read SQL file %s: %w", name, err)
	}
	return string(content), nil
}

func Migrate(db *sql.DB) error {
	entries, err := readMigrationFiles()

	if err != nil {
		return err
	}

	for _, e := range entries {

		if !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}

		sqlContent, err := readSQLFile(e.Name())

		if err != nil {
			return err
		}

		if _, err := db.Exec(sqlContent); err != nil {
			return fmt.Errorf("error executing migrations %s: %w", e.Name(), err)
		}
	}
	return nil
}
