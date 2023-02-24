package main

import (
	"database/sql"
	"log"

	"github.com/ThoPham02/simp_bank/api"
	db "github.com/ThoPham02/simp_bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/simp_bank?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/simp_bank?sslmode=disable")
	if err != nil {
		log.Fatal("can't connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Router.Run(serverAddress)
	if err != nil {
		log.Fatal("can't connect to server")
	}
}
