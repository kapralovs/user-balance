package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/user-balance/internal/balance/delivery/http"
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
	conn := postgres.NewConnection()
	pgRepo := postgres.NewRepo(conn)
	pgUsecase := usecase.New(pgRepo)
	pgHandler := http.NewHandler(pgUsecase)
	pgHandler.RegisterHTTPEndpoints(pgUsecase, s.router)

	return s.router.Listen(s.port)
}
