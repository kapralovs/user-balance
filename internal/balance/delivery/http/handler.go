package http

import (
	"encoding/json"
	"log"

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
	user, err := h.uc.GetBalanceInfo(1)
	if err != nil {
		log.Fatal(err)
	}
	json, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return c.SendString(string(json))
}
