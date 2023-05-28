package tictac

func CheckRows(
	board [][]string,
	streakToWin int,
	calculateXY func(i int) (int, int),
) bool {
	var rowStreak int
	var previousMark string
	for i := 0; i < streakToWin*2-1; i++ {
		y, x := calculateXY(i)

		if y < 0 || x < 0 || len(board) <= y || len(board[y]) <= x {
			continue
		}

		mark := board[y][x]
		if mark != "" && previousMark == mark {
			rowStreak++
		} else {
			rowStreak = 0
		}

		if rowStreak+1 == streakToWin {
			return true
		}

		previousMark = mark
	}

	return false
}

func CheckGameOver(
	board [][]string,
	cursorY int,
	cursorX int,
	streakToWin int,
) bool {
	if len(board) == 0 {
		return false
	}

	// check horizontal
	if CheckRows(
		board,
		streakToWin,
		func(i int) (int, int) {
			return cursorY, cursorX - streakToWin + 1 + i
		},
	) {
		return true
	}

	// check vertical
	if CheckRows(
		board,
		streakToWin,
		func(i int) (int, int) {
			return cursorY - streakToWin + 1 + i, cursorX
		},
	) {
		return true
	}

	// check diagonal left to right
	if CheckRows(
		board,
		streakToWin,
		func(i int) (int, int) {
			return cursorY - streakToWin + 1 + i, cursorX - streakToWin + 1 + i
		},
	) {
		return true
	}

	// check diagonal right to left
	if CheckRows(
		board,
		streakToWin,
		func(i int) (int, int) {
			return cursorY + streakToWin - 1 - i, cursorX - streakToWin + 1 + i
		},
	) {
		return true
	}

	return false
}
