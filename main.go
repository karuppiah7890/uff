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
				fmt.Printf("pass exactly two arguments\n\n")
				cli.ShowAppHelpAndExit(c, 1)
				return nil
			}

			foodFile := c.Args().Get(0)
			version := c.Args().Get(1)

			food, err := getFood(foodFile, version)
			if err != nil {
				return cli.Exit(fmt.Errorf("error occured while getting food: %v", err), 1)
			}

			fmt.Printf("existing food version: %v\n", food.Version)
			fmt.Printf("upgrading fish food %s to version %s ...\n", foodFile, version)
			return nil
		},
	}

	app.Run(os.Args)
}
