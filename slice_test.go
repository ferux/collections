package collections

import (
	"reflect"
	"testing"
)

func TestFilterInplace(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		keepFunc func(int) bool
		expected []int
	}{
		{
			name:     "Keep even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			keepFunc: func(n int) bool { return n%2 == 0 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "Keep numbers greater than 3",
			input:    []int{1, 2, 3, 4, 5, 6},
			keepFunc: func(n int) bool { return n > 3 },
			expected: []int{4, 5, 6},
		},
		{
			name:     "Keep all numbers",
			input:    []int{1, 2, 3, 4, 5},
			keepFunc: func(n int) bool { return true },
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Keep no numbers",
			input:    []int{1, 2, 3, 4, 5},
			keepFunc: func(n int) bool { return false },
			expected: []int{},
		},
		{
			name:     "Empty input",
			input:    []int{},
			keepFunc: func(n int) bool { return n > 0 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterInplace(tt.input, tt.keepFunc)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FilterInplace() = %v, want %v", result, tt.expected)
			}
		})
	}
}
