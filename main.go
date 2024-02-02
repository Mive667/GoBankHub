package main

import (
	"database/sql"
	"log"

	"github.com/Mive667/GoBankHub/api"
	db "github.com/Mive667/GoBankHub/db/sqlc"
	"github.com/Mive667/GoBankHub/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to server:", err)
	}
}
