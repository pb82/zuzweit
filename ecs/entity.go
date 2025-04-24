package ecs

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/data_structures"
)

type Entity struct {
	id uuid.UUID

	// optional
	// only certain, permanently accessible entities should be named
	name           string
	active         bool
	components     data_structures.List[Component]
	pendingRemoval data_structures.Queue[ComponentType]
	manager        *EntityManager
}

func NewEntity(manager *EntityManager) *Entity {
	return &Entity{
		id:      uuid.New(),
		active:  true,
		manager: manager,
	}
}

func (e *Entity) Id() uuid.UUID {
	return e.id
}

func (e *Entity) HasComponent(componentType ComponentType) bool {
	return e.components.Has(func(c Component) bool {
		return c.Type() == componentType
	})
}

func (e *Entity) GetComponent(componentType ComponentType) Component {
	found, component := e.components.Get(func(c Component) bool {
		return c.Type() == componentType
	})
	if found {
		return component
	}
	return nil
}

func (e *Entity) addComponent(component Component) Component {
	if e.HasComponent(component.Type()) {
		return component
	}
	e.components.Add(component)
	return component
}

func (e *Entity) removeComponent(componentType ComponentType) {
	e.components.Remove(func(c Component) bool {
		return c.Type() == componentType
	})
}

func (e *Entity) RemoveComponent(componentType ComponentType) {
	e.pendingRemoval.Push(componentType)
}

func (e *Entity) Collect() {
	for {
		if e.pendingRemoval.Empty() {
			return
		}

		next, _ := e.pendingRemoval.Pop()
		e.removeComponent(next)
	}
}

func (e *Entity) Update(dt float64) {
	e.components.ForEach(func(c Component) {
		c.Update(dt)
	})
}

func (e *Entity) Render(screen *ebiten.Image) {
	e.components.ForEach(func(c Component) {
		c.Render(screen)
	})
}

func (e *Entity) KeyInput(keys []ebiten.Key) {
	e.components.ForEach(func(c Component) {
		c.KeyInput(keys)
	})
}
