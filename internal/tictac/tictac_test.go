package tictac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFull(t *testing.T) {
	type arguments struct {
		board [][]string
	}
	tests := []struct {
		name      string
		arguments arguments
		expected  bool
	}{
		{
			name: "nil board, expect true",
			arguments: arguments{
				board: nil,
			},
			expected: true,
		},
		{
			name: "empty board, expect true",
			arguments: arguments{
				board: [][]string{},
			},
			expected: true,
		},
		{
			name: "board without empty fields, expect true",
			arguments: arguments{
				board: [][]string{{"x"}},
			},
			expected: true,
		},
		{
			name: "board with empty fields, expect false",
			arguments: arguments{
				board: [][]string{{""}},
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(
					t,
					tt.expected,
					IsFull(tt.arguments.board),
				)
			},
		)
	}
}

func TestEmptyCells(t *testing.T) {
	type arguments struct {
		board [][]string
	}
	tests := []struct {
		name      string
		arguments arguments
		expected  [][2]int
	}{
		{
			name: "nil board, expect nil",
			arguments: arguments{
				board: nil,
			},
			expected: nil,
		},
		{
			name: "empty board, expect nil",
			arguments: arguments{
				board: [][]string{},
			},
			expected: nil,
		},
		{
			name: "board without empty fields, expect nil",
			arguments: arguments{
				board: [][]string{{"x"}},
			},
			expected: nil,
		},
		{
			name: "board with 2 empty fields, expect slice of 2 arrays",
			arguments: arguments{
				board: [][]string{{"", ""}},
			},
			expected: [][2]int{{0, 0}, {0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(
					t,
					tt.expected,
					EmptyCells(tt.arguments.board),
				)
			},
		)
	}
}
