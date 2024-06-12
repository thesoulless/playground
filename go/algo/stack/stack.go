package stack

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type node[T any] struct {
	val  T
	next *node[T]
}

type Stack[T any] struct {
	Lenght int

	head *node[T]
	tail *node[T]
}

// | 3 | --> head
// | 2 |
// | 1 | --> tail
// | _ |

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	defer func() { s.Lenght++ }()

	node := node[T]{val: val}

	if s.head == nil {
		s.head = &node
		s.tail = s.head
		return
	}

	node.next = s.head
	s.head = &node
}

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var res T
		return res, ErrEmptyStack
	}

	defer func() { s.Lenght-- }()

	val := s.head.val
	s.head = s.head.next

	return val, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.head == nil {
		var res T
		return res, ErrEmptyStack
	}

	return s.head.val, nil
}
