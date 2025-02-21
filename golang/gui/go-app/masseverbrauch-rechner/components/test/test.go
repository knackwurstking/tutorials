package test

import (
	"log/slog"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Test struct {
	app.Compo
}

func (t *Test) Render() app.UI {
	slog.Debug("Rendering component", "name", "Test")
	return app.H1().Text("Test Component!")
}
