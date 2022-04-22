package main

import (
	"log"

	"fadilmuh22/go_bank/api"
	"fadilmuh22/go_bank/db"
	"fadilmuh22/go_bank/util"
)

func main() {
	server, err := api.NewServer(util.Config{}, db.Store{})
	if err != nil {
		log.Fatalf("Server cant be started %v", err)
		return
	}
	server.RunServer("0.0.0.0:8000")
}
