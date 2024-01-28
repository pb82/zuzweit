package ecs

import "github.com/hajimehoshi/ebiten/v2"

type EntityManager struct {
	entities []*Entity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{}
}

func (e *EntityManager) AddEntity() *Entity {
	_ = NewEntity()
	return nil
}

func (e *EntityManager) Collect() {
	// collect pending removal components from entities
	for _, entity := range e.entities {
		entity.Collect()
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
