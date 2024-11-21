package btree

import (
	"errors"
	"fmt"
	"strings"
)

// insert
// delete (pop)

var (
	ErrEmptyHeap = errors.New("the heap is empty")
)

type MinHeap[T ~int | ~float32] struct {
	length int
	data   []T
}

func NewMinHeap[T ~int | ~float32](size int) *MinHeap[T] {
	return &MinHeap[T]{
		length: 0,
		data:   make([]T, size),
	}
}

func (mh *MinHeap[T]) Length() int {
	return mh.length
}

func (mh *MinHeap[T]) Debug() {
	var sb strings.Builder
	sb.WriteString("heap: [")
	for i := 0; i < mh.length-1; i++ {
		sb.WriteString(fmt.Sprintf("%v, ", mh.data[i]))
	}
	sb.WriteString(fmt.Sprintf("%v]\n", mh.data[mh.length-1]))
	fmt.Println(sb.String())
}

func (mh *MinHeap[T]) Insert(val T) {
	mh.data[mh.length] = val
	mh.up(mh.length)
	mh.length++
}

func (mh *MinHeap[T]) Delete(val int) (T, int, error) {
	if mh.length == 0 {
		return 0, 0, ErrEmptyHeap
	}

	out := mh.data[0]
	mh.length--

	if mh.length == 1 {
		mh.data = nil
		mh.length--
		return out, 0, nil
	}

	mh.data[0] = mh.data[mh.length]
	mh.down(0)

	return out, 0, nil
}

func (mh *MinHeap[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (mh *MinHeap[T]) leftChild(idx int) int {
	return (idx * 2) + 1
}

func (mh *MinHeap[T]) rightChild(idx int) int {
	return (idx * 2) + 2
}

func (mh *MinHeap[T]) up(idx int) {
	if idx == 0 {
		return
	}

	pi := mh.parent(idx)
	pv := mh.data[pi]
	v := mh.data[idx]

	if pv > v {
		mh.data[idx] = pv
		mh.data[pi] = v
		mh.up(pi)
	}
}

func (mh *MinHeap[T]) down(idx int) {
	if idx >= mh.length {
		return
	}

	lIdx := mh.leftChild(idx)
	rIdx := mh.rightChild(idx)

	if lIdx >= mh.length {
		return
	}

	lv := mh.data[lIdx]
	rv := mh.data[rIdx]
	v := mh.data[idx]

	if lv > rv && v > rv {
		mh.data[idx] = rv
		mh.data[rIdx] = v
		mh.down(rIdx)
	} else if rv > lv && v > lv {
		mh.data[idx] = lv
		mh.data[lIdx] = v
		mh.down(lIdx)
	}
}

func (mh *MinHeap[T]) Min() (T, error) {
	if mh.length == 0 {
		return 0, ErrEmptyHeap
	}

	return mh.data[0], nil
}
