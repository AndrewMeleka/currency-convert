package util

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// MapToOrderedSlice converts a map to a slice of its keys, sorted in ascending order.
func MapToOrderedSlice[T constraints.Ordered](m map[T]string) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
