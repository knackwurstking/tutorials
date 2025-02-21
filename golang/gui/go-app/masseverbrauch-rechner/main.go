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
	app.Route("/", func() app.Composer {
		return &components.Test{}
	})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Masseverbrauch-Rechner",
		Description: "Ein einfacher Rechner für den Masseverbrauch für Presse 0.",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server start failed", "error", err)
		os.Exit(ErrorServerStartFailed)
	}
}
