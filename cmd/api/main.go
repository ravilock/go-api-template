package main

import (
	"github.com/ravilock/go-api-template/internal/api/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
