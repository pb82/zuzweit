package data_structures

import "errors"

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Push(t T) {
	q.data = append(q.data, t)
}

func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Pop() (T, error) {
	if len(q.data) == 0 {
		var t T
		return t, errors.New("cannot pop on empty queue")
	}

	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
