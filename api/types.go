package api

import (
	"github.com/hajimehoshi/ebiten/v2"
	mini3d "github.com/pb82/mini3d/api"
)

type Canvas []byte

type GameState struct {
}

type GameContext struct {
	Gamepads []ebiten.GamepadID
	Engine   *mini3d.Engine
	Canvas   Canvas
	Map      *Map
	Atlas    mini3d.TextureAtlas
}

// Clear clears the canvas (set all pixels to black)
func (c *Canvas) Clear() {
	for i := range *c {
		(*c)[i] = 0x0
	}
}
