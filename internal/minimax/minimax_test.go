package minimax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vehsamrak/tictac/internal/tictac"
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
					tictac.IsFull(tt.arguments.board),
				)
			},
		)
	}
}

func TestMinimax(t *testing.T) {
	type arguments struct {
		data Data
	}
	tests := []struct {
		name      string
		arguments arguments
		expected  int
	}{
		// {
		// 	name: "board with no empty cells and draw combination, must return 0",
		// 	arguments: arguments{
		// 		data: Data{
		// 			board: [][]string{
		// 				{"o", "x", "o"},
		// 				{"x", "o", "x"},
		// 				{"x", "o", "x"},
		// 			},
		// 			streakToWin: 3,
		// 			currentMark: "x",
		// 			cursorX:     1,
		// 			cursorY:     0,
		// 		},
		// 	},
		// 	expected: 0,
		// },
		// {
		// 	name: "board with no empty cells and win combination, must return 10",
		// 	arguments: arguments{
		// 		data: Data{
		// 			board: [][]string{
		// 				{"o", "x", "o"},
		// 				{"x", "x", "x"},
		// 				{"o", "o", "x"},
		// 			},
		// 			streakToWin: 3,
		// 			currentMark: "x",
		// 			cursorX:     1,
		// 			cursorY:     1,
		// 		},
		// 	},
		// 	expected: 10,
		// },
		// {
		// 	name: "board with no empty cells and lose combination, must return -10",
		// 	arguments: arguments{
		// 		data: Data{
		// 			board: [][]string{
		// 				{"o", "x", "o"},
		// 				{"x", "o", "x"},
		// 				{"x", "o", "o"},
		// 			},
		// 			streakToWin: 3,
		// 			currentMark: "x",
		// 			cursorX:     2,
		// 			cursorY:     2,
		// 		},
		// 	},
		// 	expected: -10,
		// },
		// {
		// 	name: "board with empty cells and winning combination on depth 1, must return 9",
		// 	arguments: arguments{
		// 		data: Data{
		// 			board: [][]string{
		// 				{"o", "x", "o"},
		// 				{"x", "o", "x"},
		// 				{"x", "x", ""},
		// 			},
		// 			streakToWin: 3,
		// 			currentMark: "x",
		// 			cursorX:     0,
		// 			cursorY:     2,
		// 		},
		// 	},
		// 	expected: -9,
		// },
		{
			name: "board with empty cells and winning combination on depth 1, must return 9",
			arguments: arguments{
				data: Data{
					board: [][]string{
						{"x", ""},
					},
					streakToWin: 2,
					currentMark: "x",
					cursorX:     0,
					cursorY:     0,
				},
			},
			expected: -9,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				minimax := &Minimax{}
				assert.Equal(
					t,
					tt.expected,
					minimax.Minimax(tt.arguments.data),
				)
			},
		)
	}
}
