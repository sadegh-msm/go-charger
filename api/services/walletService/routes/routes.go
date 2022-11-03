package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"wallet-service/handlers"
)

// Router creating new router and add middlewares and routes and return echo object
func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/newacc", handlers.NewAccount)
	e.POST("/increase", handlers.Increment)
	e.POST("/decrease", handlers.Decrement)
	e.GET("/balance", handlers.Balance)

	return e
}
