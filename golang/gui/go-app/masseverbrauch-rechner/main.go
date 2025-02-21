package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/MatusOllah/slogcolor"
	"github.com/SuperPaintman/nice/cli"
	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"masseverbrauch-rechner/pages"
)

const (
	ErrorServerStartFailed = 10

	Version = "v0.0.0"
)

func main() {
	cli := cli.App{
		Name:  "picow-led-server",
		Usage: cli.Usage("PicoW LED Server"),
		Action: cli.ActionFunc(func(cmd *cli.Command) cli.ActionRunner {
			var debug bool
			var generatePages bool

			_ = cli.BoolVar(cmd, &debug, "debug",
				cli.Usage("Enable debug mode"),
				cli.WithShort("d"),
				cli.Optional,
			)

			_ = cli.BoolVar(cmd, &generatePages, "generate-pages",
				cli.Usage("Generate static pages"),
				cli.WithShort("g"),
				cli.Optional,
			)

			return func(cmd *cli.Command) error {
				setupLogging()

				app.Route("/", func() app.Composer {
					return &pages.Root{}
				})

				app.RunWhenOnBrowser()

				// TODO: Add app icons
				handler := &app.Handler{
					Name:        "Masseverbrauch-Rechner",
					Description: "Ein einfacher Rechner für den Masseverbrauch für Presse 0.",
					Lang:        "de",
					Styles: []string{
						"/web/ui-v2.0.0.css",
					},
					HTML: func() app.HTMLHtml {
						return app.Html().
							Styles(map[string]string{
								"width":  "100%",
								"height": "100%",
							}).
							Attr("data-theme", "auto")
					},
					Body: func() app.HTMLBody {
						return app.Body().
							Styles(map[string]string{
								"width":  "100%",
								"height": "100%",
							})
					},
				}

				if generatePages {
					return app.GenerateStaticWebsite("dist", handler)
				}

				return serve(handler)
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

func setupLogging() {
	slogcolor.DefaultOptions.Level = slog.LevelDebug

	slog.SetDefault(
		slog.New(
			slogcolor.NewHandler(
				os.Stderr, slogcolor.DefaultOptions,
			),
		),
	)
}

func serve(handler *app.Handler) error {
	http.Handle("/", handler)

	slog.Info("Starting server", "address", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("server start failed: %w", err)
	}

	return nil
}
