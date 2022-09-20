package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/user-balance/internal/balance"
)

func RegisterHTTPEndpoints(uc balance.Usecase, router *fiber.App) {
	handler := NewHandler(uc)
	router.Get("/", handler.GetBalanceInfo)
}
