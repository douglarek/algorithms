// Package sorts provides algorithms outside the standard library for sorting
// slices and user-defined collections.
package sorts

import (
	"sort"
)

// Select sorts a *sort.Interface* using select-sort algorithms.
func Select(a sort.Interface) {
	n := a.Len()
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Swap(i, min)
	}
}

// SelectInts sorts an int array.
func SelectInts(a []int) {
	Select(sort.IntSlice(a))
}
