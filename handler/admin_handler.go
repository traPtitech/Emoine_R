package handler

import (
	"log/slog"
	"os"
)

type AdminAPIHandler struct {
	logger *slog.Logger
}

func NewAdminAPIHandler() *AdminAPIHandler {
	return &AdminAPIHandler{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})),
	}
}
