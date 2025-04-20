package api

import (
	"github.com/hajimehoshi/ebiten/v2"
	mini3d "github.com/pb82/mini3d/api"
)

type Canvas []byte

type Direction int

const (
	Forward  Direction = 0
	Left               = -90
	Right              = 90
	Backward           = 180
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
