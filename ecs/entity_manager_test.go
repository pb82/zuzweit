package ecs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EntityManager_AddEntity(t *testing.T) {
	manager := NewEntityManager()
	e1 := manager.AddEntity()
	_ = manager.AddEntity()

	assert.Equal(t, manager.entities.Size(), 0)
	manager.Collect()
	assert.Equal(t, manager.entities.Size(), 2)
	manager.RemoveEntity(e1)
	assert.Equal(t, manager.entities.Size(), 2)
	manager.Collect()
	assert.Equal(t, manager.entities.Size(), 1)
}

func Test_EntityManager_NamedEntities(t *testing.T) {
	manager := NewEntityManager()
	e1 := manager.AddNamedEntity("player")
	manager.Collect()
	assert.Equal(t, manager.entities.Size(), 1)
	assert.Equal(t, e1, manager.GetNamedEntity("player"))
	manager.RemoveEntity(e1)
	manager.Collect()
	assert.Equal(t, manager.entities.Size(), 0)
	assert.Nil(t, manager.GetNamedEntity("player"))
}
