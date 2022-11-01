package routes

import (
	"challange/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router creating new router and add middlewares and routes and return echo object
func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/newacc", handlers.NewAccount)

	return e
}
