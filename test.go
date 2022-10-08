package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
    fmt.Fprintf(os.Stdout, "Request: %+v\n", string(reqBody))
}

func main() {
    e := echo.New()

    e.HideBanner = true
    e.HidePort = true

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
        Skipper: func(c echo.Context) bool {
            if c.Request().Header.Get("X-Debug") == "" {
                return true
            }
            return false
        },
        Handler: dumpHandler,
    }))

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!\n")
    })

    e.Logger.Fatal(e.Start(":8000"))
}