package api

import (
	"github.com/gin-gonic/gin"
	"go-migration-app/config"
)

type Server struct {
	Router *gin.Engine
	Config config.Config
}

func NewServer(cfg config.Config) *Server {
	server := &Server{
		Router: gin.Default(),
		Config: cfg,
	}
	return server
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
