package ecs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EntityManager_AddEntity(t *testing.T) {
	manager := NewEntityManager()
	e1 := manager.AddEntity()
	_ = manager.AddEntity()

	assert.Len(t, manager.entities, 0)
	manager.Collect()
	assert.Len(t, manager.entities, 2)
	manager.RemoveEntity(e1)
	assert.Len(t, manager.entities, 2)
	manager.Collect()
	assert.Len(t, manager.entities, 1)
}
