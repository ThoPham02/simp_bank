package main

import (
	"database/sql"
	"log"

	"github.com/ThoPham02/simp_bank/api"
	db "github.com/ThoPham02/simp_bank/db/sqlc"
	"github.com/ThoPham02/simp_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Router.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("can't connect to server")
	}
}
