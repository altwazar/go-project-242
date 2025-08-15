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
		Usage:     "show file size",
		UsageText: "hexlet-path-size [PATH]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)"},
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
			size, err := code.GetSize(path)

			if err != nil {
				log.Fatal(err)
			}

			h := cmd.Bool("human")
			// Отформатированный размер
			fsize := code.FormatSize(size, h)
			out := fmt.Sprintf("%s\t%s", fsize, path)
			fmt.Println(out)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
