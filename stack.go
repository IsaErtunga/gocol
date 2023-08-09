package gocol

type stack[T any] struct {
	items []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		items: make([]T, 0),
	}
}

func (s *stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *stack[T]) Pop() (v T) {
	l := len(s.items)
	el := s.items[l-1]
	s.items = s.items[:l-1]
	return el
}

func (s *stack[T]) Len() int {
	return len(s.items)
}
