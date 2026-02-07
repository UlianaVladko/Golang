package main

import (
	"testing"
)

func TestSumAll(t *testing.T) {
	test := []struct {
		name     string
		numbers  []int
		expected int
	}{{
		name:     "Example 1",
		numbers:  []int{1, 2, 3},
		expected: 6,
	},
		{
			name:     "Example 2",
			numbers:  []int{10, -2, 4, 7},
			expected: 19,
		},
	}

	for _, tt  := range test{
		t.Run(tt.name, func(t *testing.T) {
			result:=sumAll(tt.numbers...)
			if result!=tt.expected{
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
