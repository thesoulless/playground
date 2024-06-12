package queue_test

import (
	"errors"
	"testing"

	"github.com/thesoulles/playground/go/algo/queue"
)

func TestQueue(t *testing.T) {
	q := queue.New[int]()

	if val, err := q.Peek(); !errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("queue is not emprty: %v", val)
		t.FailNow()
	}

	if val, err := q.Dequeue(); !errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("queue is not emprty: %v", val)
		t.FailNow()
	}

	q.Enqueue(13)

	if val, err := q.Peek(); errors.Is(err, queue.ErrEmptyQueue) || val != 13 {
		t.Errorf("Peek failed, Got: \nerr: %v, val: %v\nExpected: \nerr: nil, val: 13", err, val)
		t.FailNow()
	}

	if val, err := q.Dequeue(); errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("queue is not emprty: %v", val)
		t.FailNow()
	}

	if val, err := q.Peek(); !errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("queue is not emprty: %v", val)
		t.FailNow()
	}

	q.Enqueue(9)
	q.Enqueue(19)
	q.Enqueue(-1)

	if val, err := q.Peek(); errors.Is(err, queue.ErrEmptyQueue) || val != 9 {
		t.Errorf("Peek failed, Got: \nerr: %v, val: %v\nExpected: \nerr: nil, val: 9", err, val)
		t.FailNow()
	}

	q.Dequeue()

	if val, err := q.Peek(); errors.Is(err, queue.ErrEmptyQueue) || val != 19 {
		t.Errorf("Peek failed, Got: \nerr: %v, val: %v\nExpected: \nerr: nil, val: 19", err, val)
		t.FailNow()
	}

	q.Dequeue()
	q.Dequeue()

	if val, err := q.Dequeue(); !errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("queue is not emprty: %v", val)
		t.FailNow()
	}
}
