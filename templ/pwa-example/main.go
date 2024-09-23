package main

import (
	"log"
	"templ-pwa-example/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	greet := components.Greet("Lizzy The Cat", 2)
	handler := templ.Handler(greet)

	// log.Fatalln(http.ListenAndServe(":8080", handler))
	e.Server.Handler = handler

	if err := e.Start(":1323"); err != nil {
		log.Fatalln(err)
	}
}
