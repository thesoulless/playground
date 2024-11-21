package maps

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[int](2)

	lru.Put("1", 1)
	lru.Put("2", 2)
	v, ok := lru.Get("1")

	if !ok {
		t.Errorf("Expected %v, got %v", true, ok)
	}

	_, ok = lru.Get("3")
	if ok {
		t.Errorf("Expected %v, got %v", false, ok)
	}

	if v != 1 {
		t.Errorf("Expected %v, got %v", 1, v)
	}

	lru.Put("3", 3)
	lru.Put("4", 4)

	lru.Put("3", 33)

	v, ok = lru.Get("3")
	if !ok {
		t.Errorf("Expected %v, got %v", true, ok)
	}

	if v != 33 {
		t.Errorf("Expected %v, got %v", 33, v)
	}

	v = lru.LRU()

	if v != 4 {
		t.Errorf("Expected %v, got %v", 2, v)
	}
}
