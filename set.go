package collections

import (
	"cmp"
	"slices"
)

type Set[K comparable] map[K]struct{}

func NewSetExtract[U comparable, T any](items []T, mapper func(T) U) Set[U] {
	out := make(map[U]struct{}, len(items))
	for _, item := range items {
		out[mapper(item)] = struct{}{}
	}
	return out
}

func NewSet[K comparable, T ~[]K](items T) Set[K] {
	out := make(map[K]struct{}, len(items))
	for _, item := range items {
		out[item] = struct{}{}
	}

	return out
}

func (s Set[K]) Add(value K) {
	s[value] = struct{}{}
}

func (s Set[K]) Empty() bool {
	return len(s) == 0
}

func (s Set[K]) Contains(item K) bool {
	_, ok := s[item]
	return ok
}

// Pops pop multiple items from set
func (s Set[K]) Pops(items ...K) (subtracted []K) {
	out := make([]K, 0, len(items))
	for _, item := range items {
		if !s.Contains(item) {
			continue
		}
		out = append(out, item)
		delete(s, item)
	}
	return out
}

func (s Set[K]) Remove(item K) {
	delete(s, item)
}

type AsSliceOpt[K comparable, T ~[]K] func(T)

func SortOutput[K cmp.Ordered, T ~[]K](items T) {
	slices.Sort(items)
}

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
