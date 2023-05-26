package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkGameOver(t *testing.T) {
	const (
		height      = 3
		width       = 3
		streakToWin = 3
		X           = "X"
	)
	tests := []struct {
		name     string
		board    [][]cell
		cursorY  int
		cursorX  int
		expected bool
	}{
		{
			name:     "given empty board, expect game is not over",
			board:    [][]cell{},
			expected: false,
		},
		{
			name: "given board with 2 same marks on first line and 1 on second line, expect game is not over",
			board: [][]cell{
				{{}, {mark: X}, {mark: X}},
				{{mark: X}, {}, {}},
			},
			cursorY:  0,
			cursorX:  1,
			expected: false,
		},
		{
			name: "given board with 3 same marks horizontally, expect game is over",
			board: [][]cell{
				{{mark: X}, {mark: X}, {mark: X}},
			},
			cursorY:  0,
			cursorX:  0,
			expected: true,
		},
		{
			name: "given board with 3 same marks vertically, expect game is over",
			board: [][]cell{
				{{mark: X}, {}, {}},
				{{mark: X}, {}, {}},
				{{mark: X}, {}, {}},
			},
			cursorY:  0,
			cursorX:  0,
			expected: true,
		},
		{
			name: "given board with 3 same marks diagonally from left, expect game is over",
			board: [][]cell{
				{{mark: X}, {}, {}},
				{{}, {mark: X}, {}},
				{{}, {}, {mark: X}},
			},
			cursorY:  0,
			cursorX:  0,
			expected: true,
		},
		{
			name: "given board with 3 same marks diagonally from right, expect game is over",
			board: [][]cell{
				{{}, {}, {mark: X}},
				{{}, {mark: X}, {}},
				{{mark: X}, {}, {}},
			},
			cursorY:  0,
			cursorX:  2,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				model := NewModel(height, width, streakToWin)
				model.board = tt.board
				model.cursorY = tt.cursorY
				model.cursorX = tt.cursorX
				assert.Equal(
					t,
					tt.expected,
					model.checkGameOver(tt.cursorY, tt.cursorX),
				)
			},
		)
	}
}
