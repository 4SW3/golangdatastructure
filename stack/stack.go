package stack

import "fmt"

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = s.top.prev
	s.length--
	return n.value
}

func (s *Stack) Push(val interface{}) {
	n := &node{val, s.top}
	s.top = n
	s.length++
}

func (s *Stack) Repr() string {
	var out string
	cur := s.top

	for cur != nil {
		out += fmt.Sprintf("|%v|\n", cur.value)
		cur = cur.prev
	}

	return out
}
