package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/simp_bank?sslmode=disable")
	if err != nil {
		panic(err)
	}

	
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}


	migrate, err := migrate.NewWithDatabaseInstance(
		"file:db/migration",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	migrate.Up()
}
