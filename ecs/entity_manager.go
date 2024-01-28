package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/data_structures"
)

type EntityManager struct {
	entities []*Entity

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
	newEntities := make([]*Entity, 0)
	for _, entity := range e.entities {
		if toRemove == entity {
			continue
		}
		newEntities = append(newEntities, entity)
	}
	e.entities = newEntities
}

func (e *EntityManager) Collect() {
	// collect pending removal components from entities
	for _, entity := range e.entities {
		entity.Collect()
	}

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
		e.entities = append(e.entities, entity)
	}
}

func (e *EntityManager) Update(dt float64) {
	for _, entity := range e.entities {
		entity.Update(dt)
	}
}

func (e *EntityManager) Render(screen *ebiten.Image) {
	for _, entity := range e.entities {
		entity.Render(screen)
	}
}
