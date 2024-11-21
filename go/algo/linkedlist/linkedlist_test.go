package linkedlist_test

import (
	"testing"

	"github.com/thesoulles/playground/go/algo/linkedlist"
)

func TestLinkedList_Length(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{
			name: "no item",
			vals: nil,
			want: 0,
		},
		{
			name: "three items",
			vals: []int{5, 3, 1},
			want: 3,
		},
		{
			name: "five item",
			vals: []int{5, 3, 1, 0, 19},
			want: 5,
		},
		{
			name: "one item",
			vals: []int{5},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := linkedlist.New[int]()
			for _, v := range tt.vals {
				ll.Prepend(v)
			}

			l := ll.Length()

			if l != tt.want {
				t.Errorf("TestLinkedList_Get: %s failed. Expected: %d, Got: %d", tt.name, tt.want, l)
				return
			}

			t.Logf("TestLinkedList_Get: %s passed. Expected: %d, Got: %d",
				tt.name, tt.want, l)
		})
	}
}

func TestLinkedList_InsertAt(t *testing.T) {
	tests := []struct {
		name    string
		vals    []int
		i       int
		val     int
		want    int
		wantLen int
		wantErr bool
	}{
		{
			name:    "first",
			vals:    []int{5, 3, 1},
			i:       0,
			val:     19,
			want:    19,
			wantLen: 4,
		},
		{
			name:    "second",
			vals:    []int{5, 3, 1},
			i:       1,
			val:     19,
			want:    19,
			wantLen: 4,
		},
		{
			name:    "before last",
			vals:    []int{5, 3, 1},
			i:       2,
			val:     -1,
			want:    -1,
			wantLen: 4,
		},
		{
			name:    "last",
			vals:    []int{5, 3, 1},
			i:       3,
			val:     0,
			want:    0,
			wantLen: 4,
		},
		{
			name:    "error",
			vals:    []int{5, 3, 1},
			i:       4,
			val:     0,
			want:    0,
			wantLen: 3,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := linkedlist.New[int]()
			for _, v := range tt.vals {
				ll.Append(v)
			}

			err := ll.InsertAt(tt.i, tt.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestLinkedList_InsertAt: %s failed. Expected Err: %t, got Err: %v", tt.name, tt.wantErr, err)
				return
			}

			v, err := ll.Get(tt.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestLinkedList_InsertAt: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if v != tt.want {
				t.Errorf("TestLinkedList_InsertAt: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if ll.Length() != tt.wantLen {
				t.Errorf("TestLinkedList_InsertAt: %s failed. Expected: %d, Got: %d", tt.name, tt.wantLen, ll.Length())
				return
			}

			t.Logf("TestLinkedList_InsertAt: %s passed. Expected: %d, Got: %d. Expected Err: %t, got Err: %v",
				tt.name, tt.want, v, tt.wantErr, err)
		})
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		i        int
		want     int
		wantLen  int
		wantErr1 bool
		wantErr2 bool
	}{
		{
			name:    "first",
			vals:    []int{5, 3, 1},
			i:       0,
			want:    3,
			wantLen: 2,
		},
		{
			name:    "second",
			vals:    []int{5, 3, 1},
			i:       1,
			want:    1,
			wantLen: 2,
		},
		{
			name:     "last",
			vals:     []int{5, 3, 1},
			i:        2,
			want:     0,
			wantLen:  2,
			wantErr2: true,
		},
		{
			name:     "not found 1",
			vals:     []int{5, 3, 1},
			i:        -1,
			want:     0,
			wantLen:  3,
			wantErr1: true,
			wantErr2: true,
		},
		{
			name:     "not found 2",
			vals:     []int{5, 3, 1},
			i:        3,
			want:     0,
			wantLen:  3,
			wantErr1: true,
			wantErr2: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := linkedlist.New[int]()
			for _, v := range tt.vals {
				ll.Append(v)
			}

			err := ll.RemoveAt(tt.i)
			if (err != nil) != tt.wantErr1 {
				t.Errorf("TestLinkedList_RemoveAt: %s failed. Expected Err: %t, got Err: %v", tt.name, tt.wantErr1, err)
				return
			}

			v, err := ll.Get(tt.i)
			if (err != nil) != tt.wantErr2 {
				t.Errorf("TestLinkedList_RemoveAt: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if v != tt.want {
				t.Errorf("TestLinkedList_RemoveAt: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if ll.Length() != tt.wantLen {
				t.Errorf("TestLinkedList_RemoveAt: %s failed. Expected length: %d, Got length: %d", tt.name, tt.wantLen, ll.Length())
				return
			}

			t.Logf("TestLinkedList_RemoveAt: \"%s\" passed. Expected: %d, Got: %d. Expected Err: %t, got Err: %v",
				tt.name, tt.want, v, tt.wantErr2, err)
		})
	}

	t.Run("multiple items", func(t *testing.T) {
		vals := []int{5, 3, 1}
		ll := linkedlist.New[int]()
		for _, v := range vals {
			ll.Append(v)
		}

		for i := 0; i < len(vals); i++ {
			err := ll.RemoveAt(0)
			if err != nil {
				t.Errorf("TestLinkedList_RemoveAt: multiple items failed. Expected Err: nil, got Err: %v", err)
				return
			}
		}

		if ll.Length() != 0 {
			t.Errorf("TestLinkedList_RemoveAt: multiple items failed. Expected length: 0, Got length: %d", ll.Length())
			return
		}
	})
}

func TestLinkedList_Remove(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		i        int
		val      int
		want     int
		wantErr1 bool
		wantErr2 bool
	}{
		{
			name: "first",
			vals: []int{5, 3, 1},
			i:    0,
			val:  5,
			want: 3,
		},
		{
			name: "second",
			vals: []int{5, 3, 1},
			i:    1,
			val:  3,
			want: 1,
		},
		{
			name:     "last",
			vals:     []int{5, 3, 1},
			i:        2,
			val:      1,
			want:     0,
			wantErr2: true,
		},
		{
			name:     "not found 1",
			vals:     []int{5, 3, 1},
			val:      0,
			want:     0,
			wantErr1: true,
			wantErr2: true,
		},
		{
			name:     "not found 2",
			vals:     []int{5, 3, 1},
			val:      19,
			want:     0,
			wantErr1: true,
			wantErr2: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := linkedlist.New[int]()
			for _, v := range tt.vals {
				ll.Append(v)
			}

			err := ll.Remove(tt.val)
			if (err != nil) != tt.wantErr1 {
				t.Errorf("TestLinkedList_Remove: %s failed. Expected Err: %t, got Err: %v", tt.name, tt.wantErr1, err)
				return
			}

			if tt.wantErr1 {
				return
			}

			v, err := ll.Get(tt.i)
			if (err != nil) != tt.wantErr2 {
				t.Errorf("TestLinkedList_Remove: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if v != tt.want {
				t.Errorf("TestLinkedList_Remove: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if ll.Length() != len(tt.vals)-1 {
				t.Errorf("TestLinkedList_Remove: %s failed. Expected: %d, Got: %d", tt.name, len(tt.vals)-1, ll.Length())
				return
			}

			t.Logf("TestLinkedList_Remove: %s passed. Expected: %d, Got: %d. Expected Err: %t, got Err: %v",
				tt.name, tt.want, v, tt.wantErr2, err)
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	tests := []struct {
		name    string
		vals    []int
		i       int
		want    int
		wantErr bool
	}{
		{
			name: "first",
			vals: []int{5, 3, 1},
			i:    0,
			want: 5,
		},
		{
			name: "second",
			vals: []int{5, 3, 1},
			i:    1,
			want: 3,
		},
		{
			name: "first",
			vals: []int{5, 3, 1},
			i:    2,
			want: 1,
		},
		{
			name:    "not found",
			vals:    []int{5, 3, 1},
			i:       3,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := linkedlist.New[int]()
			for _, v := range tt.vals {
				ll.Append(v)
			}

			v, err := ll.Get(tt.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestLinkedList_Get: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if v != tt.want {
				t.Errorf("TestLinkedList_Get: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if ll.Length() != len(tt.vals) {
				t.Errorf("TestLinkedList_Get: %s failed. Expected: %d, Got: %d", tt.name, len(tt.vals), ll.Length())
				return
			}

			t.Logf("TestLinkedList_Get: %s passed. Expected: %d, Got: %d. Expected Err: %t, got Err: %v",
				tt.name, tt.want, v, tt.wantErr, err)
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	tests := []struct {
		name    string
		vals    []int
		i       int
		want    int
		wanrErr bool
	}{
		{
			name: "first",
			vals: []int{5, 3, 1},
			i:    0,
			want: 1,
		},
		{
			name: "second",
			vals: []int{5, 3, 1},
			i:    1,
			want: 3,
		},
		{
			name: "first",
			vals: []int{5, 3, 1},
			i:    2,
			want: 5,
		},
		{
			name:    "not found",
			vals:    []int{5, 3, 1},
			i:       -1,
			want:    0,
			wanrErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := linkedlist.New[int]()
			for _, v := range tt.vals {
				ll.Prepend(v)
			}

			v, err := ll.Get(tt.i)
			if (err != nil) != tt.wanrErr {
				t.Errorf("TestLinkedList_Prepend: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
			}

			if v != tt.want {
				t.Errorf("TestLinkedList_Prepend: %s failed. Expected: %d, Got: %d", tt.name, tt.want, v)
				return
			}

			if ll.Length() != len(tt.vals) {
				t.Errorf("TestLinkedList_Prepend: %s failed. Expected: %d, Got: %d", tt.name, len(tt.vals), ll.Length())
				return
			}

			t.Logf("TestLinkedList_Prepend: \"%s\" passed. Expected: %d, Got: %d. Expected Err: %t, got Err: %v",
				tt.name, tt.want, v, tt.wanrErr, err)
		})
	}
}
