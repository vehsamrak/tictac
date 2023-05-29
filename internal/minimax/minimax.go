package minimax

import (
	"github.com/vehsamrak/tictac/internal/tictac"
)

const (
	scoreWin  = 10
	scoreLose = -10
	scoreDraw = 0
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
	isGameOver := tictac.CheckGameOver(
		data.board,
		data.cursorY,
		data.cursorX,
		data.streakToWin,
	)

	if isGameOver {
		if data.board[data.cursorY][data.cursorX] == data.currentMark {
			return scoreLose
		}

		return scoreWin
	}

	if tictac.IsFull(data.board) {
		return scoreDraw
	}

	emptyCells := tictac.GetEmptyCells(data.board)
	for _, emptyCell := range emptyCells {
		emptyCellY, emptyCellX := emptyCell[0], emptyCell[1]

		predictedBoard := data.board
		predictedBoard[emptyCellY][emptyCellX] = "x"

		m.Minimax(Data{
			cursorY:     emptyCellY,
			cursorX:     emptyCellX,
			currentMark: data.currentMark,
			streakToWin: data.streakToWin,
			board:       predictedBoard,
		})
	}

	// isGameOver := tictac.CheckGameOver(
	// 	data.board,
	// 	data.cursorY,
	// 	data.cursorX,
	// 	data.streakToWin,
	// )
	//
	// if tictac.IsFull(data.board) {
	// 	if isGameOver {
	// 		if data.board[data.cursorY][data.cursorX] == data.currentMark {
	// 			return 10
	// 		}
	//
	// 		return -10
	// 	}
	//
	// 	return 0
	// }

	// collect all empty cells
	// calculate minimax for each
	// invert mark
	// add mark on board
	// take max of each minimax and return

	return 0
}
