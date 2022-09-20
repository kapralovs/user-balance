package server

import "github.com/gofiber/fiber/v2"

type server struct {
	port   string
	router *fiber.App
}

func New(port string) *server {
	return &server{
		port:   port,
		router: fiber.New(),
	}
}
