package server

import (
	"fmt"
	"courses-api-mysql-and-cb/internal/server/handler/courses"
	"courses-api-mysql-and-cb/internal/server/handler/health"
	"log"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	srv.engine.Use(cors.Default())
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/create-course", courses.CreateHandler())
	s.engine.GET("/get-courses", courses.GetHandler())
}
