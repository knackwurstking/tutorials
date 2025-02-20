package main

import (
	"strings"

	"github.com/SuperPaintman/nice/cli"
)

func main() {
	app := cli.App{
		Name:  "git",
		Usage: cli.Usage("The stupid content tracker"),
		Commands: []cli.Command{
			{
				Name:  "add",
				Usage: cli.Usage("Add file contents to the index"),
				Action: cli.ActionFunc(func(cmd *cli.Command) cli.ActionRunner {
					pathspec := cli.RestStrings(cmd, "pathspec", cli.Usage("Files to add content from"))

					return func(cmd *cli.Command) error {
						cmd.Printf("%+v\n", pathspec)

						return nil
					}
				}),
			},
			{
				Name:  "clone",
				Usage: cli.Usage("Clone a repository into a new directory"),
				Action: cli.ActionFunc(func(cmd *cli.Command) cli.ActionRunner {
					jobs := cli.Int(
						cmd, "jobs",
						cli.WithShort("j"),
						cli.WithLong("jobs"),
						cli.Usage("Specify number of submodules cloned in parallel"),
					)

					repository := cli.StringArg(
						cmd, "repository",
						cli.Usage("The (possibly remote) repository to clone from"),
					)

					directory := cli.StringArg(
						cmd, "directory",
						cli.Usage("The name of the new directory to clone into"),
						cli.Optional,
					)

					return func(cmd *cli.Command) error {
						_ = jobs
						_ = repository
						_ = directory

						if *directory == "" {
							// NOTE: from `...knackwurstking/reponame` to `reponame`
							idx := strings.LastIndex(*repository, "/")
							if idx != -1 {
								*directory = (*repository)[idx+1:]
							} else {
								*directory = *repository
							}
						}

						cmd.Printf("Cloning into '%s'...\n", *directory)

						return nil
					}
				}),
			},
			cli.CompletionCommand(),
		},
		CommandFlags: []cli.CommandFlag{
			cli.HelpCommandFlag(),
			cli.VersionCommandFlag("2.29.2"),
		},
	}

	app.HandleError(app.Run())
}
