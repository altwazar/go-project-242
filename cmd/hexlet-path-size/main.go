package main

import (
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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Нужен один аргумент - путь
			if cmd.NArg() != 1 {
				cli.ShowAppHelp(cmd)
				return cli.Exit("Error: requires one argument - path", 1)
			}

			path := cmd.Args().Get(0)
			size, err := GetSize(path)

			if err != nil {
				log.Fatal(err)
			}
			out := fmt.Sprintf("%dB \t%s", size, path)
			fmt.Println(out)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// Возвращает размер файла или файлов в директории первого уровня
func GetSize(path string) (int64, error) {
	pinfo, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}
	var size int64 = 0
	// Если директория, то перебор файлов в ней
	if pinfo.IsDir() {
		files, err := os.ReadDir(path)

		if err != nil {
			return 0, err
		}
		for _, file := range files {
			finfo, _ := file.Info()
			// Только размер файлов
			if !finfo.IsDir() {
				size = size + finfo.Size()
			}
		}
	} else {
		size = size + pinfo.Size()
	}
	return size, nil
}
