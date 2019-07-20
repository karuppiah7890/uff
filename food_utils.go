package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/fishworks/gofish"
	"github.com/spf13/afero"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"io"
	"net/http"
	"strings"
)

func getFood(foodFile string) (*gofish.Food, error) {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile(foodFile); err != nil {
		return nil, err
	}
	var food gofish.Food
	if err := gluamapper.Map(l.GetGlobal("food").(*lua.LTable), &food); err != nil {
		return nil, err
	}
	return &food, nil
}

func findUpgradedFood(foodFile, existingVersion, version string) (*gofish.Food, error) {
	fs := afero.NewOsFs()
	l := lua.NewState()
	defer l.Close()

	existingVersionFood, err := afero.ReadFile(fs, foodFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", foodFile, err)
	}
	versionUpgradedFood := strings.ReplaceAll(string(existingVersionFood), existingVersion, version)

	if err := l.DoString(versionUpgradedFood); err != nil {
		return nil, err
	}
	var food gofish.Food
	if err := gluamapper.Map(l.GetGlobal("food").(*lua.LTable), &food); err != nil {
		return nil, err
	}

	for index, foodPackage := range food.Packages {
		resp, err := http.Get(foodPackage.URL)
		if err != nil {
			return nil, fmt.Errorf("error while downloading package to calculate shasum: %v", err)
		}

		h := sha256.New()
		if _, err := io.Copy(h, resp.Body); err != nil {
			return nil, fmt.Errorf("error while calculating shasum of package: %v", err)
		}
		_ = resp.Body.Close()
		food.Packages[index].SHA256 = fmt.Sprintf("%x", h.Sum(nil))
	}

	return &food, nil
}

func upgradeFoodFile(foodFile string, food *gofish.Food, newFood *gofish.Food) error {
	fs := afero.NewOsFs()
	info, err := fs.Stat(foodFile)
	if err != nil {
		return fmt.Errorf("error finding info of file %s: %v", foodFile, err)
	}
	mode := info.Mode()
	existingVersionFood, err := afero.ReadFile(fs, foodFile)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", foodFile, err)
	}

	versionUpgradedFood := strings.ReplaceAll(string(existingVersionFood), food.Version, newFood.Version)

	packageUpgradedFood := versionUpgradedFood

	for index, foodPackage := range food.Packages {
		packageUpgradedFood = strings.ReplaceAll(packageUpgradedFood, foodPackage.SHA256, newFood.Packages[index].SHA256)
	}

	err = afero.WriteFile(fs, foodFile, []byte(packageUpgradedFood), mode)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", foodFile, err)
	}
	return nil
}
