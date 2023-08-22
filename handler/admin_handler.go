package handler

import (
	"log/slog"

	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
)

type AdminAPIHandler struct {
	logger *slog.Logger
}

func NewAdminAPIHandler(logger *slog.Logger) emoine_rv1connect.AdminAPIServiceHandler {
	return &AdminAPIHandler{
		logger: logger,
	}
}
