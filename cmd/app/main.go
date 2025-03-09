package main

import (
	"fmt"
	"social-network/config"
	"social-network/migrations"
	"social-network/pkg/database"
	"social-network/pkg/server"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.GetDbUrl())
	db := database.New(cfg.GetDbUrl())
	migrations.Migrate()
	server.NewGinServer(cfg, db).Start()
}
