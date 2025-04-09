package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Stack(t *testing.T) {
	q := Stack[int]{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	el, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, el, 3)

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el, 3)

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el, 2)

	el, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, el, 1)

	_, err = q.Pop()
	assert.Error(t, err)
}
