package handlers

import (
	"container-monitoring/internal/api/repository"
)

type BaseHandler struct {
	repo repository.Repository
}

func NewBaseHandler(r repository.Repository) *BaseHandler {
	return &BaseHandler{
		repo: r,
	}
}
