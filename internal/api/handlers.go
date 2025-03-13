package api

import "github.com/ST359/rest-api-todo/internal/service"

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc}
}
