package queue

import (
	"errors"
)

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type node[T any] struct {
	val  T
	next *node[T]
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() *Queue[T] {
	q := Queue[T]{}

	return &q
}

func (q *Queue[T]) Enqueue(val T) {
	n := &node[T]{
		val: val,
	}

	if q.tail == nil {
		q.tail = n
		/*&node[T]{
			val: val,
		}*/

		if q.head == nil {
			q.head = q.tail
		}

		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.head == nil {
		var res T
		return res, ErrEmptyQueue
	}

	h := *q.head
	q.head.next = nil
	q.head = h.next

	if h.next == nil {
		q.tail = nil
	}

	return h.val, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.head == nil {
		var res T
		return res, ErrEmptyQueue
	}

	return q.head.val, nil
}

/*func (q *Queue[T]) Visit() {
	h := q.head

	if h == nil {
		fmt.Printf("head was empty\n")
	}

	for h != nil {
		fmt.Printf("val: %v\n", h.val)
		h = h.next
	}

	fmt.Printf("head: %+v\n", q.head)
	fmt.Printf("tail: %+v\n", q.tail)
}*/
