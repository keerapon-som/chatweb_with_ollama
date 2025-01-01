package handlers

import "github.com/gofiber/fiber/v2"

type PrivateHandler struct {
}

func (h *PrivateHandler) Init(root fiber.Router) {
	root.Get("/privatetest", h.PrivateElement)
}

func (h *PrivateHandler) PrivateElement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "helloza privatetest",
	})
}
