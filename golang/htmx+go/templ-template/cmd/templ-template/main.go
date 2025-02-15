package main

import (
	"context"
	"templ-template/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DB struct {
}

func main() {
	app := echo.New()

	defaultLoggerConfig := middleware.DefaultLoggerConfig
	defaultLoggerConfig.Format = "[${time_rfc3339}] ${method} ${uri} ${status}\n"
	defaultLoggerConfig.Output = app.Logger.Output()

	app.Use(
		middleware.LoggerWithConfig(defaultLoggerConfig),
		withUser,
	)

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Static("/assets", "./assets")

	app.Start(":3000") // NOTE: dev port number 3000
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "knackwurstking@mail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
