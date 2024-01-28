package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zuzweit/ecs"
)

func Test_Queue(t *testing.T) {
	q := Queue[int]{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	el, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el, 1)

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el, 2)

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el, 3)

	_, err = q.Pop()
	assert.Error(t, err)
}

func Test_PointerQueue(t *testing.T) {
	q := Queue[*ecs.Entity]{}

	e1 := ecs.NewEntity()
	e2 := ecs.NewEntity()
	e3 := ecs.NewEntity()

	q.Push(e1)
	q.Push(e2)
	q.Push(e3)

	el, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el.Id(), e1.Id())

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el.Id(), e2.Id())

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el.Id(), e3.Id())

	_, err = q.Pop()
	assert.Error(t, err)
}
