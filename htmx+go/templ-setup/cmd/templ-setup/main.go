package main

import (
	"templ-setup/handler"

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
	)

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000") // NOTE: dev port number 3000
}
