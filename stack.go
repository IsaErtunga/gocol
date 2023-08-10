package gocol

import "errors"

var errEmpty error = errors.New("empty array")

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

func (s *stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var val T
		return val, errEmpty
	}
	l := len(s.items)
	el := s.items[l-1]
	s.items = s.items[:l-1]
	return el, nil
}

func (s *stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var val T
		return val, errEmpty
	}
	return s.items[len(s.items)-1], nil
}

func (s *stack[T]) Len() int {
	return len(s.items)
}
