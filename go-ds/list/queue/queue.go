package queue

import "errors"

type Queue[K comparable] struct {
	start  *QNode[K]
	end    *QNode[K]
	length int
}

type QNode[K comparable] struct {
	Next  *QNode[K]
	Value K
}

func newNode[K comparable](v K) *QNode[K] {
	return &QNode[K]{Value: v}
}

func NewQueue[K comparable]() *Queue[K] {
	return &Queue[K]{}
}

func (q *Queue[K]) Enqueue(v K) *Queue[K] {
	n := newNode(v)
	if q.Empty() {
		q.start = n
	} else {
		q.end.Next = n
	}
	q.end = n
	q.length++

	return q
}

func (q *Queue[K]) Dequeue() *Queue[K] {
	if !q.Empty() {
		if q.start == q.end {
			q.start, q.end = nil, nil
		} else {
			q.start = q.start.Next
		}
		q.length--
	}

	return q
}

func (q *Queue[K]) Empty() bool {
	return q.start == nil
}

func (q *Queue[K]) Each(fn func(n *QNode[K], i int)) {
	if !q.Empty() {
		c := q.start
		i := 0

		for c != nil {
			fn(c, i)
			c = c.Next
			i++
		}
	}
}

func (q *Queue[K]) First() (K, error) {
	if q.start == nil {
		return *new(K), errors.New("queue is empty")
	}

	return q.start.Value, nil
}

func (q *Queue[K]) Last() (K, error) {
	if q.start == nil {
		return *new(K), errors.New("queue is empty")
	}

	return q.end.Value, nil
}
