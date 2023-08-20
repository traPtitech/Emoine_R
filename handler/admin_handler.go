package handler

import (
	"log/slog"
)

type AdminAPIHandler struct {
	logger *slog.Logger
}

func NewAdminAPIHandler(logger *slog.Logger) *AdminAPIHandler {
	return &AdminAPIHandler{
		logger: logger,
	}
}
