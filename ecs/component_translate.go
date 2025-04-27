package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	api2 "github.com/pb82/mini3d/api"
)

type TranslateComponent struct {
	parent  *Entity
	X, Y    float64
	Advance float64
	Turn    float64
}

func NewTranslateComponent(parent *Entity, x, y float64) *TranslateComponent {
	return parent.addComponent(&TranslateComponent{
		parent:  parent,
		X:       x,
		Y:       y,
		Advance: 0,
		Turn:    0,
	}).(*TranslateComponent)
}

func (t *TranslateComponent) Render(screen *ebiten.Image) {
	// empty
}

func (t *TranslateComponent) Update(_ float64) {}

func (t *TranslateComponent) KeyInput(keys []ebiten.Key) {}

func (t *TranslateComponent) Type() ComponentType {
	return TranslateComponentType
}

func (t *TranslateComponent) Increment(x, y float64) api2.Vector3d {
	return api2.Vector3d{
		X: t.X + x,
		Y: t.Y + y,
		Z: 0,
		W: 0,
	}
}

func (t *TranslateComponent) Parent() *Entity {
	return t.parent
}
