package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/user-balance/internal/balance/repository/postgres"
	"github.com/kapralovs/user-balance/internal/balance/usecase"
)

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

func (s *server) Run() error {
	postgresRepo, err := postgres.NewPostgresRepo()
	if err != nil {
		log.Fatal(err)
	}
	postgresUsecase, err := usecase.New(postgresRepo)
	if err != nil {
		log.Fatal(err)
	}

	return s.router.Listen(s.port)
}
