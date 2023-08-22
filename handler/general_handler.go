package handler

import (
	"log/slog"

	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
)

type GeneralAPIHandler struct {
	logger *slog.Logger
}

func NewGeneralAPIHandler(logger *slog.Logger) emoine_rv1connect.GeneralAPIServiceHandler {
	return &GeneralAPIHandler{
		logger: logger,
	}
}
