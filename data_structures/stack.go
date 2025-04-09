package data_structures

import "errors"

type Stack[T any] struct {
	data []T
}

func (q *Stack[T]) Push(t T) {
	q.data = append([]T{t}, q.data...)
}

func (q *Stack[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Stack[T]) Pop() (T, error) {
	if len(q.data) == 0 {
		var t T
		return t, errors.New("cannot pop on empty stack")
	}

	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}

func (q *Stack[T]) Peek() (T, error) {
	if q.Empty() {
		var t T
		return t, errors.New("cannot peek on empty stack")
	}

	return q.data[0], nil
}
