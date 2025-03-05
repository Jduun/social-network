package server

import (
	"fmt"
	"net/http"
	"social-network/config"
	"social-network/pkg/database"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	engine *gin.Engine
	db     *database.PostgresDatabase
	config *config.Config
}

func NewGinServer(config *config.Config, db *database.PostgresDatabase) Server {
	engine := gin.Default()

	return &ginServer{
		engine: engine,
		db:     db,
		config: config,
	}
}

func (server *ginServer) Start() {
	server.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := server.engine.Run(fmt.Sprintf(":%s", server.config.HTTPServer.Port))
	if err != nil {
		panic(fmt.Sprintf("Cannot run Gin Server: %s", err))
	}
}
