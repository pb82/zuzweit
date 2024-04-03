package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/data_structures"
)

type EntityManager struct {
	entities      data_structures.List[*Entity]
	pendingAdd    data_structures.Queue[*Entity]
	pendingRemove data_structures.Queue[*Entity]
}

func NewEntityManager() *EntityManager {
	return &EntityManager{}
}

func (e *EntityManager) AddEntity() *Entity {
	entity := NewEntity()
	e.pendingAdd.Push(entity)
	return entity
}

func (e *EntityManager) RemoveEntity(entity *Entity) {
	e.pendingRemove.Push(entity)
}

func (e *EntityManager) removeEntity(toRemove *Entity) {
	e.entities.Remove(func(e *Entity) bool {
		return e == toRemove
	})
}

func (e *EntityManager) Collect() {
	// collect pending removal components from entities
	e.entities.ForEach(func(e *Entity) {
		e.Collect()
	})

	// collect pending removal entities
	for {
		if e.pendingRemove.Empty() {
			break
		}

		entity, _ := e.pendingRemove.Pop()
		e.removeEntity(entity)
	}

	// insert pending entities
	for {
		if e.pendingAdd.Empty() {
			break
		}

		entity, _ := e.pendingAdd.Pop()
		e.entities.Add(entity)
	}
}

func (e *EntityManager) Update(dt float64) {
	e.entities.ForEach(func(e *Entity) {
		e.Update(dt)
	})
}

func (e *EntityManager) Render(screen *ebiten.Image) {
	e.entities.ForEach(func(e *Entity) {
		e.Render(screen)
	})
}
