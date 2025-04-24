package api

import (
	"github.com/hajimehoshi/ebiten/v2"
	mini3d "github.com/pb82/mini3d/api"
)

type Canvas []byte

type Direction int

const (
	North Direction = 0
	West            = -90
	East            = 90
	South           = 180
)

func (d Direction) Angle() float64 {
	return float64(d)
}

type GameState struct {
}

type GameContext struct {
	Gamepads []ebiten.GamepadID
	Engine   *mini3d.Engine
	Canvas   Canvas
}

// Clear clears the canvas (set all pixels to black)
func (c *Canvas) Clear() {
	for i := range *c {
		(*c)[i] = 0x0
	}
}
