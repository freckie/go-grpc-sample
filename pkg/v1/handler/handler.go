package handler

import "frec.kr/tdoo/pkg/v1/service"

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Task() *TaskHandler {
	return NewTaskHandler(h.svc.Task())
}

func (h *Handler) User() *UserHandler {
	return NewUserHandler(h.svc.User())
}
