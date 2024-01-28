package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
