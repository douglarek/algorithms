package sorts

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type aint = sort.IntSlice

func TestSelect(t *testing.T) {
	type args struct {
		a sort.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{"a", args{aint([]int{5, 2, 4, 0, 1})}},
		{"b", args{aint([]int{0, 1, 2, 4, 3})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.args.a
			Select(a)
			assert.Equal(t, true, sort.IsSorted(a))
		})
	}
}

func TestSelectInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"a", args{[]int{5, 2, 4, 0, 1}}},
		{"b", args{[]int{0, 1, 2, 4, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.args.a
			SelectInts(a)
			assert.Equal(t, true, sort.IsSorted(aint(a)))
		})
	}
}

func benchmarkSelect(a aint, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select(a)
	}
}

func BenchmarkSelect1(b *testing.B) { benchmarkSelect([]int{10, 0, 1, 3, 2, 4, 5}, b) }
func BenchmarkSelect2(b *testing.B) { benchmarkSelect([]int{-1, 0, 11, 3, 2, 74, 5}, b) }
