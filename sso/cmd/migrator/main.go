package main

import (
	"errors"
	"flag"
	"fmt"

	_ "database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	_ "gopkg.in/check.v1" // возможно не надо
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "storage path")
	flag.StringVar(&migrationsPath, "migrations-path", "", "migrations path")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name og migrations table")
	flag.Parse()
	if storagePath == "" {
		panic("storage path is required")
	}

	if migrationsPath == "" {
		panic("migrations path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)

	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrate to apply")
			return
		}
		panic(err)
	}

}
