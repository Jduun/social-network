package main

import (
	"social-network/config"
	"social-network/pkg/database"
	"social-network/pkg/server"
)

func main() {
	config := config.MustLoad()
	db := database.New("")
	server.NewGinServer(config, db).Start()
}
