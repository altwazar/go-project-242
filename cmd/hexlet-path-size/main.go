package main

import (
	"code"
	"context"
	"fmt"
	"github.com/urfave/cli/v3" // imports as package "cli"
	"log"
	"os"
)

func main() {
	// (&cli.Command{}).Run(context.Background(), os.Args)

	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		UsageText: "hexlet-path-size [global options]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories"},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)"},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Нужен один аргумент - путь
			if cmd.NArg() != 1 {
				err := cli.ShowAppHelp(cmd)

				if err != nil {
					log.Fatal(err)
				}
				return cli.Exit("Error: requires one argument - path", 1)
			}

			path := cmd.Args().Get(0)
			all := cmd.Bool("all")
			recursive := cmd.Bool("recursive")
			size, err := code.GetSize(path, all, recursive)

			if err != nil {
				log.Fatal(err)
			}

			human := cmd.Bool("human")
			// Отформатированный размер
			fsize := code.FormatSize(size, human)
			out := fmt.Sprintf("%s\t%s", fsize, path)
			fmt.Println(out)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
