package linkedlist

import (
	"cmp"
	"errors"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrNodeNotFound    = errors.New("node not found")
)

type Node[T cmp.Ordered] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
}

type LinkedList[T cmp.Ordered] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func New[T cmp.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Length() int {
	return ll.length
}

func (ll *LinkedList[T]) InsertAt(index int, val T) error {
	if index < 0 || index > ll.length {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		ll.Prepend(val)
		return nil
	}

	if index == ll.length {
		ll.Append(val)
		return nil
	}

	node := &Node[T]{
		Value: val,
	}

	p := ll.head

	for i := 0; i < index; i++ {
		p = p.next
	}

	ll.length++
	node.next = p
	node.prev = p.prev

	if p.prev != nil {
		p.prev.next = node
	}
	p.prev = node

	return nil
}

func (ll *LinkedList[T]) RemoveAt(i int) error {
	if i < 0 || i >= ll.length {
		return ErrIndexOutOfRange
	}

	p := ll.head

	for j := 0; j < i; j++ {
		p = p.next
	}

	if p.prev != nil {
		p.prev.next = p.next
	}

	if p.next != nil {
		p.next.prev = p.prev
	}

	p = p.next

	if i == 0 {
		ll.head = p
	}

	if i == ll.length-1 {
		ll.tail = p
	}

	ll.length--

	return nil
}

func (ll *LinkedList[T]) RemoveNode(node *Node[T]) error {
	if node == nil {
		return ErrNodeNotFound
	}

	ll.length--

	if node == ll.head {
		ll.head = node.next
	}

	if node == ll.tail {
		ll.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	return nil
}

func (ll *LinkedList[T]) Remove(val T) error {
	h := ll.head

	for h != nil && h.Value != val {
		h = h.next
	}

	if h == nil {
		return ErrNodeNotFound
	}

	if h.Value != val {
		return ErrNodeNotFound
	}

	ll.length--

	if h == ll.head {
		ll.head = h.next
	}

	if h == ll.tail {
		ll.tail = h.prev
	}

	if h.prev != nil {
		h.prev.next = h.next
	}

	if h.next != nil {
		h.next.prev = h.prev
	}

	return nil
}

func (ll *LinkedList[T]) Append(val T) {
	node := &Node[T]{
		Value: val,
	}

	ll.length++

	if ll.tail == nil {
		ll.head = node
		ll.tail = node
		return
	}

	tail := ll.tail
	ll.tail = node
	ll.tail.prev = tail
	tail.next = node
}

func (ll *LinkedList[T]) Prepend(val T) {
	node := &Node[T]{
		Value: val,
	}

	ll.length++

	if ll.head == nil {
		ll.head = node
		ll.tail = node
		return
	}

	head := ll.head
	ll.head = node
	head.prev = ll.head
	ll.head.next = head
}

func (ll *LinkedList[T]) Get(i int) (T, error) {
	if i < 0 || i >= ll.length {
		var res T
		return res, ErrIndexOutOfRange
	}

	res := ll.head
	for j := 0; j < i; j++ {
		res = res.next
	}

	return res.Value, nil
}

func (ll *LinkedList[T]) GetNode(i int) (*Node[T], error) {
	if i < 0 || i >= ll.length {
		return nil, ErrIndexOutOfRange
	}

	res := ll.head
	for j := 0; j < i; j++ {
		res = res.next
	}

	return res, nil
}
