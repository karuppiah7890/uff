package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kyokomi/emoji"
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

			food, err := getFood(foodFile)
			if err != nil {
				return cli.Exit(fmt.Errorf("error occured while getting food: %v", err), 1)
			}

			existingVersion := food.Version
			fmt.Printf("existing food version: %v\n", existingVersion)
			if existingVersion == version {
				fmt.Printf("fish food %s is already in version %s. nothing to upgrade!\n", foodFile, version)
				return nil
			}
			fmt.Printf("upgrading fish food %s to version %s ...\n", foodFile, version)

			newFood, err := findUpgradedFood(foodFile, existingVersion, version)
			if err != nil {
				return cli.Exit(fmt.Errorf("error while finding upgraded food: %v", err), 1)
			}

			err = upgradeFoodFile(foodFile, food, newFood)
			if err != nil {
				return cli.Exit(fmt.Errorf("error while upgrading food file: %v", err), 1)
			}

			done := emoji.Sprint("done! :tropical_fish:")
			fmt.Println(done)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("error occurred while running the command: %v", err)
	}
}
