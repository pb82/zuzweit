package ecs

import "github.com/hajimehoshi/ebiten/v2"

type ComponentType uint32

const (
	TranslateComponentType ComponentType = iota
	MenuComponentType
)

type Component interface {
	Type() ComponentType
	Update(dt float64)
	Render(screen *ebiten.Image)
	Parent() *Entity
}
