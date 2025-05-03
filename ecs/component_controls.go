package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	mini3d "github.com/pb82/mini3d/api"
	"zuzweit/api"
)

type ControlsComponent struct {
	parent *Entity
	engine *mini3d.Engine
}

func NewControlsComponent(parent *Entity, engine *mini3d.Engine) *ControlsComponent {
	return parent.addComponent(&ControlsComponent{
		parent: parent,
		engine: engine,
	}).(*ControlsComponent)
}

func (t *ControlsComponent) Render(screen *ebiten.Image) {
	// empty
}

func (t *ControlsComponent) Update(_ float64) {

}

func (t *ControlsComponent) Type() ComponentType {
	return ControlsComponentType
}

func (t *ControlsComponent) Parent() *Entity {
	return t.parent
}

func (t *ControlsComponent) KeyInput(keys []ebiten.Key) {
	for _, key := range keys {
		switch key {
		case ebiten.KeyUp:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewAdvance(t.parent, 1, t.engine))
			}
			break
		case ebiten.KeyDown:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewAdvance(t.parent, -1, t.engine))
			}
			break
		case ebiten.KeyLeft:
			if api.GetCommandQueue().Empty() {
			}
			break
		case ebiten.KeyRight:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewTurnLeft(t.parent, t.engine))
			}
			break
		}
	}
}
