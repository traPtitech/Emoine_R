package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/traPtitech/Emoine_R/handler"
	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	adminAPIHandler := handler.NewAdminAPIHandler(logger)
	generalAPIHandler := handler.NewGeneralAPIHandler(logger)

	mux.Handle(emoine_rv1connect.NewAdminAPIServiceHandler(
		adminAPIHandler,
		connect.WithInterceptors(), // TODO: 権限者認証
	))
	mux.Handle(emoine_rv1connect.NewGeneralAPIServiceHandler(
		generalAPIHandler,
		connect.WithInterceptors(), // TODO: 部員or招待者認証
	))

	logger.Info("Server started")
	http.ListenAndServe(
		"localhost:8090",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
