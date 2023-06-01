package minimax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimax(t *testing.T) {
	type arguments struct {
		board [][]string
		data  Data
		depth int
	}
	tests := []struct {
		name          string
		arguments     arguments
		expectedScore int
		expectedY     int
		expectedX     int
	}{
		{
			name: "board with empty cells and winning combination on depth 1, must return 9",
			arguments: arguments{
				data: Data{
					Players:       []string{"x", "o"},
					StreakToWin:   2,
					MaximizerMark: "x",
					CursorX:       0,
					CursorY:       0,
				},
				board: [][]string{
					{"x", ""},
				},
			},
			expectedScore: 9,
			expectedY:     0,
			expectedX:     1,
		},
		{
			name: "board with empty cells and winning combination on depth 1, must return 9",
			arguments: arguments{
				board: [][]string{
					{"x", "x", "o"},
					{"x", "o", "o"},
					{"", "x", ""},
				},
				data: Data{
					Players:       []string{"x", "o"},
					StreakToWin:   3,
					MaximizerMark: "x",
					CursorX:       0,
					CursorY:       0,
				},
			},
			expectedScore: 9,
			expectedY:     2,
			expectedX:     0,
		},
		{
			name: "board with empty cells and losing combination on depth 2, must return -8",
			arguments: arguments{
				board: [][]string{
					{"o", "", ""},
					{"", "o", ""},
					{"o", "x", "x"},
				},
				data: Data{
					Players:       []string{"x", "o"},
					MaximizerMark: "x",
					StreakToWin:   3,
					CursorX:       0,
					CursorY:       0,
				},
			},
			expectedScore: -8,
			expectedY:     1,
			expectedX:     0,
		},
		{
			name: "board with empty cells more then calculating depth, must return draw prediction",
			arguments: arguments{
				board: [][]string{
					{"", "", ""},
					{"", "x", ""},
					{"", "x", "o"},
				},
				data: Data{
					Players:       []string{"o", "x"},
					MaximizerMark: "o",
					StreakToWin:   3,
					CursorY:       1,
					CursorX:       1,
				},
			},
			expectedScore: 0,
			expectedY:     0,
			expectedX:     1,
		},
		{
			name: "board with Oes turn, must make prediction to avoid loss",
			arguments: arguments{
				board: [][]string{
					{"x", "", ""},
					{"x", "x", "."},
					{"o", "x", "o"},
				},
				data: Data{
					Players:       []string{"o", "x"},
					MaximizerMark: "o",
					StreakToWin:   3,
					CursorY:       0,
					CursorX:       0,
				},
			},
			expectedScore: 0,
			expectedY:     0,
			expectedX:     1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				score, y, x := Minimax{SortResult: true}.
					Minimax(
						tt.arguments.data,
						tt.arguments.board,
						tt.arguments.depth,
					)

				assert.Equal(t, tt.expectedY, y, "y is not expected")
				assert.Equal(t, tt.expectedX, x, "x is not expected")
				assert.Equal(t, tt.expectedScore, score, "score is not expected")
			},
		)
	}
}
