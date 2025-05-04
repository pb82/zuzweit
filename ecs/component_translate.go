package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/api"
)

type TranslateComponent struct {
	parent  *Entity
	X, Y    float64
	Compass *api.Compass
}

func NewTranslateComponent(parent *Entity, x, y float64) *TranslateComponent {
	return parent.addComponent(&TranslateComponent{
		parent:  parent,
		X:       x,
		Y:       y,
		Compass: api.NewCompass(),
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

func (t *TranslateComponent) NextPosition(backwards bool) (float64, float64) {
	switch t.Compass.GetDirection() {
	case api.North:
		if !backwards {
			return t.X, t.Y + 1
		} else {
			return t.X, t.Y - 1
		}
	case api.South:
		if !backwards {
			return t.X, t.Y - 1
		} else {
			return t.X, t.Y + 1
		}
	case api.East:
		if !backwards {
			return t.X - 1, t.Y
		} else {
			return t.X + 1, t.Y
		}
	case api.West:
		if !backwards {
			return t.X + 1, t.Y
		} else {
			return t.X - 1, t.Y
		}
	default:
		return -1, -1
	}
}

func (t *TranslateComponent) Parent() *Entity {
	return t.parent
}
