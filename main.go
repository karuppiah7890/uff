package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name: "uff",
		Usage: "update (go)fish food",
		UsageText: "uff <food-file> <version>",
		Action: func(c *cli.Context) error {
			fmt.Println("upgrading fish food ...")
			return nil
		},
	}

	app.Run(os.Args)
}