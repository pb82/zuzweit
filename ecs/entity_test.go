package ecs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Entity(t *testing.T) {
	e := NewEntity()
	c := NewTranslateComponent(e)

	assert.True(t, e.HasComponent(TranslateComponentType))
	assert.Equal(t, c.Parent(), e)
}

func Test_EntityRemoveComponent(t *testing.T) {
	e := NewEntity()
	_ = NewTranslateComponent(e)

	assert.Len(t, e.components, 1)
	e.RemoveComponent(TranslateComponentType)
	assert.Len(t, e.components, 1)
	e.Collect()
	assert.Len(t, e.components, 0)
}
