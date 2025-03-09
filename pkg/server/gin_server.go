package server

import (
	"fmt"
	"log"
	"net/http"

	"social-network/config"
	"social-network/internal/handlers"
	"social-network/internal/repositories"
	"social-network/internal/services"
	"social-network/pkg/database"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	engine *gin.Engine
	db     *database.PostgresDatabase
	cfg    *config.Config
}

func NewGinServer(cfg *config.Config, db *database.PostgresDatabase) Server {
	engine := gin.Default()

	return &ginServer{
		engine: engine,
		db:     db,
		cfg:    cfg,
	}
}

func (s *ginServer) Start() {
	userRepo := repositories.NewUserPostgresRepository(s.db)
	authService := services.NewAuthServiceImpl(userRepo)
	authHandlers := handlers.NewAuthHTTPHandlers(authService)

	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	s.engine.POST("/auth/register", authHandlers.Register)
	s.engine.POST("/auth/login", authHandlers.Login)
	s.engine.GET("/auth/me", authHandlers.GetMe)

	err := s.engine.Run(fmt.Sprintf(":%s", s.cfg.AppPort))
	if err != nil {
		log.Fatalf("Cannot run Gin Server: %s", err)
	}
}
