package handlers

import "github.com/gofiber/fiber/v2"

type PublicHandler struct {
}

func (h *PublicHandler) Init(root fiber.Router) {
	root.Get("/publictest", h.PublicElement)
}

func (h *PublicHandler) PublicElement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "publictest",
	})
}
