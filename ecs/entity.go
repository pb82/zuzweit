package ecs

import "github.com/google/uuid"

type Entity struct {
	id         uuid.UUID
	active     bool
	components []Component

	pendingRemoval []ComponentType
}

func NewEntity() *Entity {
	return &Entity{
		id:     uuid.New(),
		active: true,
	}
}

func (e *Entity) Id() uuid.UUID {
	return e.id
}

func (e *Entity) HasComponent(componentType ComponentType) bool {
	for _, component := range e.components {
		if component.Type() == componentType {
			return true
		}
	}
	return false
}

func (e *Entity) RemoveComponent(componentType ComponentType) {
	e.components = make([]Component, len(e.components)-1)
	for _, component := range e.components {
		if component.Type() == componentType {
			continue
		}
		e.components = append(e.components, component)
	}
}

func (e *Entity) Update(dt float64) {
	for _, component := range e.components {
		component.Update(dt)
	}
}
