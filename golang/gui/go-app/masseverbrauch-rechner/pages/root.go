package pages

import (
	"log/slog"

	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"masseverbrauch-rechner/components"
)

type Root struct {
	app.Compo
}

func (r *Root) Render() app.UI {
	slog.Debug("Rendering page", "name", "Root")

	calc := &components.Calculator{} // TODO: Forget about this calc component

	// TODO: Create form inputs needed for calculation

	return app.Div().
		Class("ui-debug").
		Styles(map[string]string{
			"width":  "100%",
			"height": "100%",
		}).
		Body(calc.Render())
}
