package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traPtitech/Emoine_R/handler"
	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// TODO: Interceptor
	mux := http.NewServeMux()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	adminAPIHandler := handler.NewAdminAPIHandler(logger)
	generalAPIHandler := handler.NewGeneralAPIHandler(logger)

	mux.Handle(emoine_rv1connect.NewAdminAPIServiceHandler(adminAPIHandler))
	mux.Handle(emoine_rv1connect.NewGeneralAPIServiceHandler(generalAPIHandler))

	logger.Info("Server started")
	err := http.ListenAndServe(
		"localhost:8090",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		panic(err)
	}

	// TODO: 認証
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// デバッグ用 /debugを叩くと認証したものとみなす
	e.GET("/debug", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["userid"] = "aaa"
		err := sess.Save(c.Request(), c.Response())
		if err != nil {
			return c.String(http.StatusInternalServerError, "セッションの保存に失敗しました")
		}

		return c.String(http.StatusOK, "あなたの名前を"+fmt.Sprint(sess.Values["userid"])+"として認証しました")
	})

	e.POST("/oauth/generate/code", handler.OAuthGenerateCodeHandler)
	e.POST("/oauth/callback", handler.OAuthCallbackHandler)

	withLogin := e.Group("")
	withLogin.Use(handler.CheckLogin)

	withAdmin := withLogin.Group("")
	withAdmin.Use(handler.CheckIsAdmin)
}
