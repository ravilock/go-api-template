package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: uuid.NewString,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintln("OK"))
	})
	e.Logger.Fatal(e.Start(":6969"))
}
