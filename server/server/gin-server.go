package server

import (
	"fmt"
	"log"
	"witty-server/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	app    *gin.Engine
	config *config.AppConfig
}

func NewServer(config *config.AppConfig) *Server {

	ginApp := gin.New()

	gin.SetMode(config.AppMode)

	ginApp.Use(gin.Logger())
	ginApp.Use(gin.Recovery())

	return &Server{
		app:    ginApp,
		config: config,
	}
}

func (s *Server) Start() {
	serverAddr := fmt.Sprintf(":%s", s.config.ServerPort)
	log.Fatal(s.app.Run(serverAddr))

}
