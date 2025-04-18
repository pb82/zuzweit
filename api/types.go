package api

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	mini3d "github.com/pb82/mini3d/api"
)

type Canvas []byte

type Direction uint8

const (
	Forward Direction = iota
	Backward
	Left
	Right
)

type GameState struct {
}

type GameContext struct {
	Map      *tiled.Map
	Gamepads []ebiten.GamepadID
	Engine   *mini3d.Engine
	Canvas   Canvas
}

func (g *GameContext) LoadMap(fileName string) error {
	m, err := tiled.LoadFile(fileName)
	if err != nil {
		return err
	}
	g.Map = m
	return nil
}

// Clear clears the canvas (set all pixels to black)
func (c *Canvas) Clear() {
	for i := range *c {
		(*c)[i] = 0x0
	}
}
