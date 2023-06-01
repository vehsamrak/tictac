package minimax

import (
	"math"

	"github.com/vehsamrak/tictac/internal/tictac"
)

const (
	scoreWin  = 10
	scoreLose = -10
	scoreDraw = 0
	maxDepth  = 10
)

type Data struct {
	MaximizerMark string   // mark of current player
	Players       []string // players marks
	CursorY       int      // last move Y
	CursorX       int      // last move X
	StreakToWin   int      // streak of marks needed to win
}

type Minimax struct{}

// Minimax applies best move prediction algorithm to tictactoe board
func (m Minimax) Minimax(data Data, board [][]string, depth int) (score int, y int, x int) {
	isGameOver := tictac.CheckGameOver(
		board,
		data.CursorY,
		data.CursorX,
		data.StreakToWin,
	)

	if isGameOver {
		if board[data.CursorY][data.CursorX] == data.MaximizerMark {
			return scoreWin - depth, data.CursorY, data.CursorX
		}

		return scoreLose + depth, data.CursorY, data.CursorX
	}

	if tictac.IsFull(board) {
		return scoreDraw, data.CursorY, data.CursorX
	}

	if depth == maxDepth {
		return 0, data.CursorY, data.CursorX
	}

	currentMark := data.Players[0]

	isMaximizer := data.MaximizerMark == currentMark
	var value int
	if isMaximizer {
		value = math.MinInt
	} else {
		value = math.MaxInt
	}

	var predictedY int
	var predictedX int
	emptyCells := tictac.EmptyCells(board)
	for _, emptyCell := range emptyCells {
		emptyCellY, emptyCellX := emptyCell[0], emptyCell[1]

		// copying board
		predictedBoard := make([][]string, len(board))
		for i := range board {
			predictedBoard[i] = make([]string, len(board[i]))
			copy(predictedBoard[i], board[i])
		}

		// populating copied board with prediction
		predictedBoard[emptyCellY][emptyCellX] = currentMark

		minimaxValue, _, _ := m.Minimax(
			Data{
				CursorY:       emptyCellY,
				CursorX:       emptyCellX,
				MaximizerMark: data.MaximizerMark,
				StreakToWin:   data.StreakToWin,
				// switching players
				Players: append(data.Players[1:], currentMark),
			},
			predictedBoard,
			depth+1,
		)

		if isMaximizer {
			if minimaxValue > value {
				value = minimaxValue
				predictedY = emptyCellY
				predictedX = emptyCellX
			}
		} else {
			if minimaxValue < value {
				value = minimaxValue
				predictedY = emptyCellY
				predictedX = emptyCellX
			}
		}
	}

	return value, predictedY, predictedX
}
