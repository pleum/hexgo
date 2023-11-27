package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/service/game"
)

type Handler struct {
	GameService game.Service
}

func (h Handler) CreateGame() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		players := []domain.Player{{Name: "A"}, {Name: "B"}}
		g, err := h.GameService.CreateNew(players)
		if err != nil {
			return c.
				Status(http.StatusInternalServerError).
				JSON(&fiber.Map{
					"success": true,
					"error":   err.Error(),
				})
		}

		return c.
			Status(http.StatusOK).
			JSON(&fiber.Map{
				"success": true,
				"data":    g,
			})
	}
}

func NewHandler(gameService game.Service) Handler {
	return Handler{
		GameService: gameService,
	}
}
