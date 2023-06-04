package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	if q.Len() != 5 {
		t.Errorf("Length should be 5")
	}

	if q.Peek() != 1 {
		t.Errorf("Front should be 1")
	}

	if q.Dequeue() != 1 {
		t.Errorf("Dequeue should be 1")
	}

	if q.Peek() != 2 {
		t.Errorf("Front should be 2")
	}

	if q.Len() != 4 {
		t.Errorf("Length should be 4")
	}

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	if q.Peek() != nil {
		t.Errorf("Front should be nil")
	}

	if q.Dequeue() != nil {
		t.Errorf("Dequeue should be nil")
	}
}
