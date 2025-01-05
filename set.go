package collections

import (
	"cmp"
	"slices"
)

type empty struct{}

// Set is a typed map with keys only. It does not stores any values and only keys.
type Set[K comparable] map[K]empty

// NewSetExtract maps incoming slice types into another type prior to building a set.
func NewSetExtract[U comparable, T any](items []T, mapper func(T) U) Set[U] {
	out := make(map[U]empty, len(items))
	for _, item := range items {
		out[mapper(item)] = empty{}
	}
	return out
}

// NewSet converts incoming slice into set.
func NewSet[K comparable, T ~[]K](items T) Set[K] {
	out := make(map[K]empty, len(items))
	for _, item := range items {
		out[item] = empty{}
	}

	return out
}

// Add element to set.
func (s Set[K]) Add(value K) {
	s[value] = empty{}
}

// Empty checks if set contains at least one value.
func (s Set[K]) Empty() bool {
	return len(s) == 0
}

// Contains check if set contains specific item.
func (s Set[K]) Contains(item K) bool {
	_, ok := s[item]
	return ok
}

// Pops pop multiple items from set returning them back to the caller.
func (s Set[K]) Pops(items ...K) (subtracted []K) {
	out := make([]K, 0, len(items))
	for _, item := range items {
		if !s.Contains(item) {
			continue
		}
		out = append(out, item)
		s.Remove(item)
	}
	return out
}

// Remove item from set. It does not checks the existance of item before deleting.
func (s Set[K]) Remove(item K) {
	delete(s, item)
}

// AsSliceOpt allows to adjust behaviour of the AsSlice method.
type AsSliceOpt[K comparable, T ~[]K] func(T)

// SortOutput sorts output items after converting set to slice.
func SortOutput[K cmp.Ordered, T ~[]K](items T) {
	slices.Sort(items)
}

// AsSlice converts set to slice.
func (s Set[K]) AsSlice(opts ...AsSliceOpt[K, []K]) []K {
	out := make([]K, 0, len(s))
	for item := range s {
		out = append(out, item)
	}

	for _, opt := range opts {
		opt(out)
	}

	return out
}
