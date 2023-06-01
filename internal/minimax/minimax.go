package minimax

import (
	"math"
	"sort"

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

type prediction struct {
	value int
	y     int
	x     int
}

type Minimax struct {
	SortResult bool // sort result prediction to be determined, useful for tests
}

// Minimax applies best move prediction algorithm to tictactoe board
func (m Minimax) Minimax(
	data Data,
	board [][]string,
	depth int,
) (score int, y int, x int) {
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

	emptyCells := tictac.EmptyCells(board)
	values := make(chan prediction)
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

		go func() {
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

			values <- prediction{value: minimaxValue, y: emptyCellY, x: emptyCellX}
		}()
	}

	emptyCellsCount := len(emptyCells)

	var predictions []prediction
	for prediction := range values {
		predictions = append(predictions, prediction)

		emptyCellsCount--
		if emptyCellsCount == 0 {
			break
		}
	}

	if m.SortResult {
		sort.SliceStable(predictions, func(i, j int) bool {
			return predictions[i].y < predictions[j].y
		})
		sort.SliceStable(predictions, func(i, j int) bool {
			return predictions[i].x < predictions[j].x
		})
	}

	isMaximizer := data.MaximizerMark == currentMark

	var value int
	if isMaximizer {
		value = math.MinInt
	} else {
		value = math.MaxInt
	}

	var predictedY int
	var predictedX int
	for _, prediction := range predictions {
		if isMaximizer {
			if prediction.value > value {
				value = prediction.value
				predictedY = prediction.y
				predictedX = prediction.x
			}
		} else {
			if prediction.value < value {
				value = prediction.value
				predictedY = prediction.y
				predictedX = prediction.x
			}
		}
	}

	return value, predictedY, predictedX
}
