package queue

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type node[T any] struct {
	val  T
	next *node[T]
}

type Queue[T any] struct {
	len  int
	head *node[T]
	tail *node[T]
}

func New[T any]() *Queue[T] {
	q := Queue[T]{}

	return &q
}

func (q *Queue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for n := q.head; n != nil; n = n.next {
		sb.WriteString(fmt.Sprintf("%v", n.val))
		if n.next != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (q *Queue[T]) Enqueue(val T) {
	q.len++

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

	q.len--

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

func (q *Queue[T]) Length() int {
	return q.len
}
