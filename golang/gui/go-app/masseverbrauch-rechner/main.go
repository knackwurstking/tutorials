package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/MatusOllah/slogcolor"
	"github.com/SuperPaintman/nice/cli"
	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"masseverbrauch-rechner/components"
)

const (
	ErrorServerStartFailed = 10

	Version = "v0.0.0"
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
	cli := cli.App{
		Name:  "picow-led-server",
		Usage: cli.Usage("PicoW LED Server"),
		Action: cli.ActionFunc(func(cmd *cli.Command) cli.ActionRunner {
			// TODO: Flags: "--generate-pages"

			return func(cmd *cli.Command) error {
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

				// TODO: Check for flag: "--generate-pages"
				if err := generateStaticWebsite(handler); err != nil {
					return err
				}

				if err := serve(handler); err != nil {
					return err
				}

				return nil
			}
		}),
		// Commands: []cli.Command{
		// 	cli.CompletionCommand(),
		// },
		CommandFlags: []cli.CommandFlag{
			cli.HelpCommandFlag(),
			cli.VersionCommandFlag(Version),
		},
	}

	cli.HandleError(cli.Run())
}

func generateStaticWebsite(handler *app.Handler) error {
	return fmt.Errorf("not implemented") // TODO: Implement
}

func serve(handler *app.Handler) error {
	http.Handle("/", handler)

	slog.Info("Starting server", "address", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("server start failed: %w", err)
	}

	return nil
}
