package ecs

type ComponentType uint32

const (
	TranslateComponentType ComponentType = iota
)

type Component interface {
	Type() ComponentType
	Update(dt float64)
	Parent() *Entity
}
