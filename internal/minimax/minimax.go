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
	currentMark string     // mark of current player
	board       [][]string // tic-tac-toe board
	cursorY     int        // last move Y
	cursorX     int        // last move X
	streakToWin int        // streak of marks needed to win
	depth       int        // depth of current node in minmax tree
}

type Minimax struct{}

// Minimax applies best move prediction algorithm to tictactoe board
func (m *Minimax) Minimax(data Data) int {
	if tictac.IsFull(data.board) {
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
