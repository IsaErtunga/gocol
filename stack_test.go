package gocol

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	_, err := s.Pop()
	if err == nil {
		t.Fail()
	}

	_, err = s.Peek()
	if err == nil {
		t.Fail()
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Fail()
	}
	item, err := s.Pop()
	if item != 3 || err != nil {
		t.Fail()
	}

	item, err = s.Peek()
	if item != 2 || err != nil {
		t.Fail()
	}
}
