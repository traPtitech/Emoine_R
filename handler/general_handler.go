package handler

import (
	"log/slog"
)

type GeneralAPIHandler struct {
	logger *slog.Logger
}

func NewGeneralAPIHandler(logger *slog.Logger) *GeneralAPIHandler {
	return &GeneralAPIHandler{
		logger: logger,
	}
}
