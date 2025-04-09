package api

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	mini3d "github.com/pb82/mini3d/api"
)

type GameState struct {
}

type GameContext struct {
	Map      *tiled.Map
	Gamepads []ebiten.GamepadID
	Engine   *mini3d.Engine
}

func (g *GameContext) LoadMap(fileName string) error {
	m, err := tiled.LoadFile(fileName)
	if err != nil {
		return err
	}
	g.Map = m
	return nil
}
