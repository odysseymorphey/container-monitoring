package server

import (
	"container-monitoring/internal/api/repository"
	"container-monitoring/internal/api/server/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	baseHandler *handlers.BaseHandler
}

func NewRouter(repo repository.Repository) *Router {
	return &Router{baseHandler: handlers.NewBaseHandler(repo)}
}

func (router *Router) RegisterRoutes(srv *fiber.App) {
	api := srv.Group("/api")
	api.Use(cors.New())

	{
		api.Get("/get_statuses", router.baseHandler.GetStatuses)
		api.Post("/add_status", router.baseHandler.AddStatus)
	}
}
