package handler

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
}

func (h IndexHandler) HandleTemplate(ctx echo.Context) error {
	templ, _ := template.New("").ParseFiles("templates/index.html")
	return templ.ExecuteTemplate(ctx.Response().Writer, "Base", nil)
}
