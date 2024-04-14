package api

import "github.com/lafriks/go-tiled"

type GameState struct {
}

type GameContext struct {
	Map *tiled.Map
}

func (g *GameContext) LoadMap(fileName string) error {
	m, err := tiled.LoadFile(fileName)
	if err != nil {
		return err
	}
	g.Map = m
	return nil
}
