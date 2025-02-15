package main

import (
	"htmx-go-todo/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	defaultLoggerConfig := middleware.DefaultLoggerConfig
	defaultLoggerConfig.Format = "[${time_rfc3339}] ${method} ${uri} ${status}\n"
	defaultLoggerConfig.Output = app.Logger.Output()

	app.Use(
		middleware.LoggerWithConfig(defaultLoggerConfig),
	)

	indexHandler := handler.IndexHandler{}
	app.GET("/", indexHandler.HandleTemplate)

	app.Static("/assets", "./assets")

	if err := app.Start(":3000"); err != nil {
		app.Logger.Panicf("App start failed: %s", err)
	}
}
