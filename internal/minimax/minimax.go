package minimax

import (
	"github.com/vehsamrak/tictac/internal/tictac"
)

const (
	scoreWin  = 10
	scoreLose = -10
)

type Data struct {
	evaluate    func() int
	currentMark string
	board       [][]string
	cursorY     int
	cursorX     int
	streakToWin int
	depth       int
}

type Minimax struct{}

func (m *Minimax) Minimax(data Data) int {
	// if win = 10, lose = -10, draw = 0, max depth = 0
	// for each empty cell run minimax
	// if minimax()

	if isFull(data.board) {
		isGameOver := tictac.CheckGameOver(
			data.board,
			data.cursorY,
			data.cursorX,
			data.streakToWin,
		)

		if isGameOver {
			if data.board[data.cursorY][data.cursorX] == data.currentMark {
				return 10
			}

			return -10
		}

		return 0
	}

	return 0
}

// isFull checks if board has at least one empty field
func isFull(board [][]string) bool {
	for _, row := range board {
		for _, mark := range row {
			if mark == "" {
				return false
			}
		}
	}

	return true
}
