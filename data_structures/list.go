package data_structures

type List[T comparable] struct {
	items []T
}

func (l *List[T]) Empty() bool {
	return len(l.items) == 0
}

func (l *List[T]) Size() int {
	return len(l.items)
}

func (l *List[T]) Add(t T) {
	l.items = append(l.items, t)
}

func (l *List[T]) Remove(pred func(t T) bool) {
	var newData []T
	for _, item := range l.items {
		if pred(item) {
			continue
		}

		newData = append(newData, item)
	}
	l.items = newData
}

func (l *List[T]) Has(pred func(t T) bool) bool {
	for _, item := range l.items {
		if pred(item) {
			return true
		}
	}
	return false
}

func (l *List[T]) ForEach(cb func(t T)) {
	for _, item := range l.items {
		cb(item)
	}
}
