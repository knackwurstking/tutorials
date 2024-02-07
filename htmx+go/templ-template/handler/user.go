package handler

import (
	"templ-template/model"
	"templ-template/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(ctx echo.Context) error {
	return render(ctx, user.Show(model.User{
		Email: "knackwurstking@email.com",
	}))
}
