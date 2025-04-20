package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/api"
)

type TranslateComponent struct {
	parent    *Entity
	X, Y      float64
	Direction api.Direction
}

func NewTranslateComponent(parent *Entity) *TranslateComponent {
	return parent.addComponent(&TranslateComponent{
		parent:    parent,
		X:         0,
		Y:         0,
		Direction: api.Forward,
	}).(*TranslateComponent)
}

func (t *TranslateComponent) Render(screen *ebiten.Image) {
	// empty
}

func (t *TranslateComponent) Update(_ float64) {}

func (t *TranslateComponent) KeyInput(keys []ebiten.Key) {}

func (t *TranslateComponent) Type() ComponentType {
	return ControlsComponentType
}

func (t *TranslateComponent) ToCameraPosition() (x float64, y float64, z float64, yaw float64) {
	yaw = t.Direction.Angle()
	x = t.X
	y = 0.5
	z = t.Y
	return
}

func (t *TranslateComponent) Parent() *Entity {
	return t.parent
}
