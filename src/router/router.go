package router

import (
	"net/http"

	"github.com/gdsc-ys/fluentify-server/src/handler"
	userMiddleware "github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Debug = true

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/GetUser", handler.GetUser)
	e.POST("/UpdateUser", userMiddleware.AuthMiddleware(handler.UpdateUser))

	return e
}
