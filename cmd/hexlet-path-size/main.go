package main

import (
	"context"
	"github.com/urfave/cli/v3" // imports as package "cli"
	"os"
)

func main() {
	(&cli.Command{}).Run(context.Background(), os.Args)
}
