package collections

import "golang.org/x/exp/constraints"

func ForEach[S ~[]E, E any](items S, f func(E)) {
	for _, item := range items {
		f(item)
	}
}

func Reduce[S ~[]E, E any, U constraints.Ordered](items S, f func(E) U) (result U) {
	var accumulator U
	for _, item := range items {
		accumulator += f(item)
	}

	return accumulator
}

func GroupByUniqueKey[T comparable, U any](items []U, mapper func(U) T) map[T]U {
	out := make(map[T]U, len(items))
	for _, item := range items {
		out[mapper(item)] = item
	}
	return out
}

func IteratorFromSlice[T any](values []T) iterator[T] {
	return iterator[T]{
		values: values,
	}
}

type iterator[T any] struct {
	values []T
	idx    int
}

func (s *iterator[T]) Reset(newValues []T) {
	s.values = newValues
	s.idx = 0
}

func (s *iterator[T]) Size() int {
	return len(s.values)
}

func (s *iterator[T]) PickNext() (T, bool) {
	if s.idx == len(s.values) {
		var t T
		return t, false
	}

	value := s.values[s.idx]
	s.idx++
	return value, true
}

func (s *iterator[T]) Pick() T {
	if s.idx >= len(s.values) {
		var t T
		return t
	}

	value := s.values[s.idx]
	s.idx++
	return value
}

// FilterInplace filters incmoming slice by using provided func an adjusts origin slice.
func FilterInplace[S ~[]E, E any](items S, keepFunc func(E) bool) S {
	curr, next := 0, 0
	for next < len(items) {
		item := items[next]
		if keepFunc(item) {
			if curr != next {
				items[curr], items[next] = items[next], items[curr]
			}

			curr++
			next++

			continue
		}

		next++
	}

	return items[:curr]
}

func FindFirstMatch[S ~[]E, E any](items S, match func(E) bool) int {
	for i, item := range items {
		if match(item) {
			return i
		}
	}
	return -1
}

func Map[T, U any](items []T, f func(T) U) []U {
	out := make([]U, len(items))
	for i, item := range items {
		out[i] = f(item)
	}
	return out
}
