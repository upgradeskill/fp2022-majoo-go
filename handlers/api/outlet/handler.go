package outlet

import (
	"mini-pos/core/outlet/service"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
