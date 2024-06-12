package arraylist

import "errors"

var (
	ErrArrayListEmpty = errors.New("array list is empty")
)

type ArrayList[T any] struct {
	len int
	cap int
	a   []T
}

func New[T any](size int) *ArrayList[T] {
	return &ArrayList[T]{
		len: 0,
		cap: size,
		a:   make([]T, 0, size),
	}
}

func (a *ArrayList[T]) Push(value T) {
	if a.len < a.cap {
		a.a[a.len] = value
		a.len++

		return
	}

	c := make([]T, a.cap)
	copy(c, a.a)

	a.cap *= 2
	a.a = make([]T, a.cap*2)
	copy(a.a, c)

	a.a[a.len] = value
	a.len++
}

func (a *ArrayList[T]) Pop() (T, error) {
	if a.len == 0 {
		var res T
		return res, ErrArrayListEmpty
	}

	a.len--
	return a.a[a.len], nil
}

func (a *ArrayList[T]) Len() int {
	return a.len
}

func (a *ArrayList[T]) Cap() int {
	return a.cap
}

func (a *ArrayList[T]) Enqueue(index int, value T) {
	panic("not implemented")
}

func (a *ArrayList[T]) Dequeue(index int) (T, error) {
	panic("not implemented")
}
