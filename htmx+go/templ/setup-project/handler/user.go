package handler

import (
	"templ-setup/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(ctx echo.Context) error {
	return render(ctx, user.Show())
}
