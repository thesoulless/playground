package sort_test

import (
	stdsort "sort"
	"testing"

	"github.com/thesoulles/playground/go/algo/sort"
)

func TestQuick(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "zero elements",
			args: args{s: []int{}},
		},
		{
			name: "one element",
			args: args{s: []int{1}},
		},
		{
			name: "five elements",
			args: args{s: []int{5, 2, 1, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := make([]int, len(tt.args.s))
			_ = copy(sc, tt.args.s)
			sort.Quick(tt.args.s)

			stdsort.Ints(sc)

			for i := range tt.args.s {
				if tt.args.s[i] != sc[i] {
					t.Errorf("sorted arrays are not equal. got: \n%v\nExpected: \n%v\n", tt.args.s, sc)
					t.FailNow()
				}
			}
		})
	}
}
