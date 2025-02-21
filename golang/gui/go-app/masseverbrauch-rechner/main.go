package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/MatusOllah/slogcolor"
	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"masseverbrauch-rechner/components"
)

const (
	ErrorServerStartFailed = 10
)

func init() {
	// TODO: Add flags parser for "--generate-pages"
	// LINK:  https://go-app.dev/github-deploy#:~:text=GitHub%20Pages.-,GENERATE%20A%20STATIC%20WEBSITE,-A%20static%20website

	slogcolor.DefaultOptions.Level = slog.LevelDebug

	slog.SetDefault(
		slog.New(
			slogcolor.NewHandler(
				os.Stderr, slogcolor.DefaultOptions,
			),
		),
	)
}

func main() {
	app.Html().Attr("data-theme", "dark")

	app.Route("/", func() app.Composer {
		return &components.Test{}
	})

	app.RunWhenOnBrowser()

	handler := &app.Handler{
		Name:        "Masseverbrauch-Rechner",
		Description: "Ein einfacher Rechner für den Masseverbrauch für Presse 0.",
		Lang:        "de",
		Styles: []string{
			"/web/ui-v2.0.0.css",
		},
		HTML: func() app.HTMLHtml {
			return app.Html().Attr("data-theme", "auto")
		},
	}

	http.Handle("/", handler)

	slog.Info("Starting server", "address", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server start failed", "error", err)
		os.Exit(ErrorServerStartFailed)
	}
}
