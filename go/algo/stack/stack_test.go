package stack_test

import (
	"errors"
	"testing"

	"github.com/thesoulles/playground/go/algo/stack"
)

func TestStack(t *testing.T) {
	ps := stack.New[int]()

	if _, err := ps.Pop(); !errors.Is(err, stack.ErrEmptyStack) {
		t.Error("Expected error, got nil")
	}

	if _, err := ps.Peek(); !errors.Is(err, stack.ErrEmptyStack) {
		t.Error("Expected error, got nil")
	}

	ps.Push(1)
	ps.Push(2)
	ps.Push(3)

	if val, err := ps.Pop(); val != 3 || err != nil {
		t.Error("Expected 3, got", val)
	}

	if val, err := ps.Peek(); val != 2 || err != nil {
		t.Error("Expected 2, got", val)
	}

	if val, err := ps.Pop(); val != 2 || err != nil {
		t.Error("Expected 2, got", val)
	}

	if val, err := ps.Pop(); val != 1 || err != nil {
		t.Error("Expected 1, got", val)
	}

	if _, err := ps.Pop(); !errors.Is(err, stack.ErrEmptyStack) {
		t.Error("Expected error, got nil")
	}
}
