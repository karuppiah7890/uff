package main

import (
	"github.com/fishworks/gofish"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func getFood(foodFile string, version string) (*gofish.Food, error) {
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
