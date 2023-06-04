package queue

import "fmt"

type (
	Queue struct {
		start, end *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	n := q.start
	q.start = n.next
	q.length--
	return n.value
}

func (q *Queue) Enqueue(val interface{}) {
	n := &node{val, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

func (q *Queue) Len() int {
	return q.length
}

func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}

func (q *Queue) Repr() {
	var out string
	cur := q.start
	for cur != nil {
		if cur.next == nil {
			out += fmt.Sprintf("%v", cur.value)
			break
		}
		out += fmt.Sprintf("%v -> ", cur.value)
		cur = cur.next
	}
}
