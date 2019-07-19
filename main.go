package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:      "uff",
		Usage:     "update (go)fish food",
		UsageText: "uff <food-file> <version>",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				fmt.Println("pass exactly two arguments\n")
				cli.ShowAppHelpAndExit(c, 1)
				return nil
			}
			foodFile := c.Args().Get(0)
			version := c.Args().Get(1)
			fmt.Printf("upgrading fish food %s to version %s ...\n", foodFile, version)
			return nil
		},
	}

	app.Run(os.Args)
}
