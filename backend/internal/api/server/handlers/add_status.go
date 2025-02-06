package handlers

import (
	"container-monitoring/internal/api/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *BaseHandler) AddStatus(c *fiber.Ctx) error {
	var status models.PingStatus
	if err := json.Unmarshal(c.Body(), &status); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.repo.AddStatus(c.Context(), &status); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}
