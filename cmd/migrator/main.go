package main

import (
	"flag"
	"fmt"

	"errors"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationPath, migrationsTable string

	flag.StringVar(&storagePath, "storage", "", "Path to the storage file")
	flag.StringVar(&migrationPath, "migrations", "", "Path to the migrations file")
	flag.StringVar(&migrationsTable, "table", "", "Name of the migrations table")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is required")
	}
	if migrationPath == "" {
		panic("migrations path is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No new migrations to apply")

			return
		} else {
			panic(err)
		}

	}
	fmt.Println("Migrations applied successfully")

}
