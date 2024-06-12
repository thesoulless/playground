package ringbuffer

import "errors"

type RingBuffer[T any] struct {
}

func New[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{}
}

func (r *RingBuffer[T]) Push(value T) {

}

func (r *RingBuffer[T]) Pop() (T, error) {
	var res T
	return res, errors.New("not implemented")
}

func (r *RingBuffer[T]) Len() int {
	return 0
}

func (r *RingBuffer[T]) Enqueue(value T) {
	panic("not implemented")
}

func (r *RingBuffer[T]) Dequeue() (T, error) {
	var res T

	return res, errors.New("not implemented")
}

func (r *RingBuffer[T]) Cap() int {
	panic("not implemented")
}
