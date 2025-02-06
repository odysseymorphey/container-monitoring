package handlers

import "github.com/gofiber/fiber/v2"

func (h *BaseHandler) GetStatuses(c *fiber.Ctx) error {
	pingStatuses, err := h.repo.GetStatuses(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(pingStatuses)
}
