package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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
	pgRepo := postgres.NewPostgresRepo()
	conn, err := pgx.Connect(context.Background(), pgRepo.Url)
	if err != nil {
		return err
	}
	pgRepo.Conn = conn
	pgUsecase := usecase.New(pgRepo)
	pgHandler := http.NewHandler(pgUsecase)
	pgHandler.RegisterHTTPEndpoints(pgUsecase, s.router)

	return s.router.Listen(s.port)
}
