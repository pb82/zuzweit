package ecs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Entity(t *testing.T) {
	em := NewEntityManager()
	e := em.AddEntity()
	c := NewTranslateComponent(e)

	assert.True(t, e.HasComponent(TranslateComponentType))
	assert.Equal(t, c.Parent(), e)
}

func Test_EntityRemoveComponent(t *testing.T) {
	em := NewEntityManager()
	e := em.AddEntity()
	_ = NewTranslateComponent(e)

	assert.Equal(t, e.components.Size(), 1)
	e.RemoveComponent(TranslateComponentType)
	assert.Equal(t, e.components.Size(), 1)
	e.Collect()
	assert.Equal(t, e.components.Size(), 0)
}
