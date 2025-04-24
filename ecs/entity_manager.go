package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zuzweit/data_structures"
)

type EntityManager struct {
	entities      data_structures.List[*Entity]
	pendingAdd    data_structures.Queue[*Entity]
	pendingRemove data_structures.Queue[*Entity]
	namedEntities map[string]*Entity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		namedEntities: make(map[string]*Entity),
	}
}

func (e *EntityManager) AddEntity() *Entity {
	entity := NewEntity(e)
	e.pendingAdd.Push(entity)
	return entity
}

func (e *EntityManager) AddNamedEntity(name string) *Entity {
	entity := NewEntity(e)
	entity.name = name
	e.pendingAdd.Push(entity)
	return entity
}

func (e *EntityManager) GetNamedEntity(name string) *Entity {
	entity := e.namedEntities[name]
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

		// if the entity is named, deregister it
		if entity.name != "" {
			delete(e.namedEntities, entity.name)
		}
	}

	// insert pending entities
	for {
		if e.pendingAdd.Empty() {
			break
		}

		entity, _ := e.pendingAdd.Pop()
		e.entities.Add(entity)

		// if the entity is named, register it
		if entity.name != "" && e.namedEntities[entity.name] == nil {
			e.namedEntities[entity.name] = entity
		}
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

func (e *EntityManager) KeyInput(keys []ebiten.Key) {
	e.entities.ForEach(func(e *Entity) {
		e.KeyInput(keys)
	})
}
