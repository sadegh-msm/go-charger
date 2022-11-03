package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"offering-service/handlers"
)

// Router creating new router and add middlewares and routes and return echo object
func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/newCode", handlers.NewCode)
	e.POST("/redeem", handlers.Redeem)
	e.GET("/codeusers", handlers.CodeUsers)

	return e
}
