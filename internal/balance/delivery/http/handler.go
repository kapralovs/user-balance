package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/user-balance/internal/balance"
)

type Handler struct {
	uc balance.Usecase
}

func NewHandler(uc balance.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) GetBalanceInfo(c *fiber.Ctx) error {
	return c.SendString("Some string")
}
