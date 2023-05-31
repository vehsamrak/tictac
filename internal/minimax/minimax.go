package minimax

import (
	"math"

	"github.com/vehsamrak/tictac/internal/tictac"
)

const (
	scoreWin  = 10
	scoreLose = -10
	scoreDraw = 0
)

type Data struct {
	maximizerMark string   // mark of current player
	Players       []string // players marks
	cursorY       int      // last move Y
	cursorX       int      // last move X
	streakToWin   int      // streak of marks needed to win
}

type Minimax struct{}

// Minimax applies best move prediction algorithm to tictactoe board
func (m *Minimax) Minimax(data Data, board [][]string, depth int) (score int, y int, x int) {
	// TODO[petr]: Remove this check
	if depth == 10 {
		panic("Maximum depth reached")
	}

	isGameOver := tictac.CheckGameOver(
		board,
		data.cursorY,
		data.cursorX,
		data.streakToWin,
	)

	if isGameOver {
		if board[data.cursorY][data.cursorX] == data.maximizerMark {
			return scoreWin - depth, data.cursorY, data.cursorX
		}

		return scoreLose + depth, data.cursorY, data.cursorX
	}

	if tictac.IsFull(board) {
		return scoreDraw, data.cursorY, data.cursorX
	}

	// switching players
	currentMark := data.Players[0]
	data.Players = data.Players[1:]
	data.Players = append(data.Players, currentMark)

	isMaximizer := data.maximizerMark == currentMark
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

		minimaxValue, turnY, turnX := m.Minimax(
			Data{
				cursorY:       emptyCellY,
				cursorX:       emptyCellX,
				maximizerMark: data.maximizerMark,
				streakToWin:   data.streakToWin,
				Players:       data.Players,
			},
			predictedBoard,
			depth+1,
		)

		if isMaximizer {
			if minimaxValue > value {
				value = minimaxValue
				predictedY = turnY
				predictedX = turnX
			}
		} else {
			if minimaxValue < value {
				value = minimaxValue
				predictedY = turnY
				predictedX = turnX
			}
		}
	}

	return value, predictedY, predictedX
}
