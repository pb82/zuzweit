package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/api"
)

type ControlsComponent struct {
	parent *Entity
}

func NewControlsComponent(parent *Entity) *ControlsComponent {
	return parent.addComponent(&ControlsComponent{
		parent: parent,
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
				api.GetCommandQueue().Push(NewAdvance(t.parent, 1))
			}
			break
		case ebiten.KeyDown:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewAdvance(t.parent, -1))
			}
			break
		case ebiten.KeyLeft:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewTurn(t.parent, -1))
			}
			break
		case ebiten.KeyRight:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewTurn(t.parent, 1))
			}
			break
		}
	}
}
