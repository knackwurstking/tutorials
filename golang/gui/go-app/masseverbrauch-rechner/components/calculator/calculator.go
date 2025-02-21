package calculator

import (
	"log/slog"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TODO: Implement the Calculator component
type Calculator struct {
	app.Compo
}

func (c *Calculator) Render() app.UI {
	slog.Debug("Rendering component", "name", "Calculator")

	return app.H1().
		Class("ui-outline-text").
		Styles(map[string]string{
			"font-size": "2.5rem",
		}).
		Text("Calculator")
}
