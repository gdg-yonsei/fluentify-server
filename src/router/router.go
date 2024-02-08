package router

import (
	"net/http"

	"github.com/gdsc-ys/fluentify-server/src/handler"
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
	e.PATCH("/UpdateUser", handler.UpdateUser)
	e.DELETE("/DeleteUser", handler.DeleteUser)

	return e
}
