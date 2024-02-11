package handler

import (
	"htmx-go-todo/view/index"

	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
}

func (h IndexHandler) HandleTemplate(ctx echo.Context) error {
	return render(ctx, index.Index())
}
