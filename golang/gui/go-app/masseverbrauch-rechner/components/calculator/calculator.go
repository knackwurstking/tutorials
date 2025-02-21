package calculator

import (
	"log/slog"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Calculator struct {
	app.Compo
}

func (c *Calculator) Render() app.UI {
	slog.Debug("Rendering component", "name", "Calculator")

	return app.Div()
}
