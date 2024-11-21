package maps

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/thesoulles/playground/go/algo/linkedlist"
)

type LRU[T cmp.Ordered] struct {
	m   map[string]*linkedlist.Node[T]
	k   map[*linkedlist.Node[T]]string
	l   *linkedlist.LinkedList[T]
	cap int
}

func NewLRU[T cmp.Ordered](size int) *LRU[T] {
	return &LRU[T]{
		cap: size,
		m:   make(map[string]*linkedlist.Node[T]),
		k:   make(map[*linkedlist.Node[T]]string),
		l:   linkedlist.New[T](),
	}
}

func (lru *LRU[T]) Get(key string) (T, bool) {
	if v, ok := lru.m[key]; ok {
		lru.l.RemoveNode(v)
		lru.l.Prepend(v.Value)
		return v.Value, true
	}

	var zero T
	return zero, false
}

func (lru *LRU[T]) LRU() T {
	v, err := lru.l.Get(lru.l.Length() - 1)
	if err != nil {
		panic(err)
	}

	return v
}

func (lru *LRU[T]) Put(key string, value T) {
	if v, ok := lru.m[key]; ok {
		lru.l.RemoveNode(v)
		lru.l.Prepend(value)
		v, err := lru.l.GetNode(0)
		if err != nil {
			panic(err)
		}

		lru.m[key] = v
		lru.k[v] = key

		return
	}

	if lru.l.Length() == lru.cap {
		v, err := lru.l.GetNode(lru.cap - 1)
		if err != nil {
			panic(err)
		}

		k := lru.k[v]

		delete(lru.k, v)
		delete(lru.m, k)

		lru.l.RemoveAt(lru.cap - 1)
	}

	lru.l.Prepend(value)
	v, err := lru.l.GetNode(0)
	if err != nil {
		panic(err)
	}

	lru.k[v] = key
	lru.m[key] = v
}

func (lru *LRU[T]) Debug() {
	var sb strings.Builder
	sb.WriteString("LRU: [")
	for i := 0; i < lru.l.Length()-1; i++ {
		v, err := lru.l.GetNode(i)
		if err != nil {
			panic(err)
		}

		sb.WriteString(fmt.Sprintf("%v", v.Value))
		sb.WriteString(", ")
	}
	v, err := lru.l.GetNode(lru.l.Length() - 1)
	if err != nil {
		panic(err)
	}

	sb.WriteString(fmt.Sprintf("%v]\n", v.Value))
	sb.WriteString("]\n")
	println(sb.String())
}
