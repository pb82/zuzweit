package ecs

import (
	"github.com/google/uuid"
	"zuzweit/data_structures"
)

type Entity struct {
	id         uuid.UUID
	active     bool
	components []Component

	pendingRemoval []data_structures.Queue[Component]
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

func (e *Entity) addComponent(component Component) Component {
	if e.HasComponent(component.Type()) {
		return component
	}
	e.components = append(e.components, component)
	return component
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
