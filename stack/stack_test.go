package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	if s.Len() != 5 {
		t.Errorf("Length should be 5")
	}

	if s.Peek() != 5 {
		t.Errorf("Top should be 5")
	}

	if s.Pop() != 5 {
		t.Errorf("Pop should be 5")
	}

	if s.Peek() != 4 {
		t.Errorf("Top should be 4")
	}

	if s.Len() != 4 {
		t.Errorf("Length should be 4")
	}

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()

	if s.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	if s.Peek() != nil {
		t.Errorf("Top should be nil")
	}

	if s.Pop() != nil {
		t.Errorf("Pop should be nil")
	}
}
