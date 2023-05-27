package client

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

type cell struct {
	mark string
}

type Player struct {
	mark string
}

type model struct {
	debugMessage    string // string with any message to render in debug mode
	message         string
	board           [][]cell
	players         []Player
	debug           bool
	gameOver        bool
	turnsLeft       int
	currentPlayerId int // id is players index
	cursorX         int
	cursorY         int
	streakToWin     int // marks streak needed to win
}

func NewModel(height, width, streakToWin int) model {
	board := make([][]cell, height)
	for y := 0; y < height; y++ {
		board[y] = make([]cell, width)
		for x := 0; x < width; x++ {
			board[y][x] = cell{}
		}
	}

	marks := []string{"X", "O"}
	var usedColors []colorful.Color
	var players []Player
	for _, mark := range marks {
		// selecting different colors for each player
		color := colorful.HappyColor()
		for _, usedColor := range usedColors {
			for {
				color = colorful.HappyColor()
				if color.DistanceLab(usedColor) >= 1 {
					break
				}
			}
		}

		usedColors = append(usedColors, color)

		players = append(players, Player{
			mark: lipgloss.NewStyle().Bold(true).Foreground(
				lipgloss.Color(color.Hex()),
			).Render(mark),
		})
	}

	currentPlayerId := 0

	return model{
		board:           board,
		cursorX:         width / 2,
		cursorY:         height / 2,
		players:         players,
		currentPlayerId: currentPlayerId,
		streakToWin:     streakToWin,
		turnsLeft:       height * width,
		message: fmt.Sprintf(
			"New game started! Goal is to occupy %d in a row. Current turn: %s",
			streakToWin,
			players[currentPlayerId].mark,
		),
	}
}

func (m *model) nextTurn(y int, x int) {
	if m.checkGameOver(y, x) {
		m.message = fmt.Sprintf("Game over. Winner is %s!", m.players[m.currentPlayerId].mark)
		m.gameOver = true
		return
	}

	m.turnsLeft--
	if m.turnsLeft == 0 {
		m.message = fmt.Sprintf("Game over. Draw!")
		m.gameOver = true
		return
	}

	if m.currentPlayerId == len(m.players)-1 {
		m.currentPlayerId = 0
	} else {
		m.currentPlayerId++
	}

	m.message = fmt.Sprintf("Current turn: %s", m.players[m.currentPlayerId].mark)
}

func (m *model) checkRows(calculateXY func(i int) (int, int)) bool {
	var rowStreak int
	var previousMark string
	for i := 0; i < m.streakToWin*2-1; i++ {
		y, x := calculateXY(i)

		if y < 0 || x < 0 || len(m.board) <= y || len(m.board[y]) <= x {
			continue
		}

		mark := m.board[y][x].mark
		if mark != "" && previousMark == mark {
			rowStreak++
		} else {
			rowStreak = 0
		}

		if rowStreak+1 == m.streakToWin {
			return true
		}

		previousMark = mark
	}

	return false
}

func (m *model) checkGameOver(cursorY int, cursorX int) bool {
	if len(m.board) == 0 {
		return false
	}

	// check horizontal
	if m.checkRows(
		func(i int) (int, int) {
			return m.cursorY, m.cursorX - m.streakToWin + 1 + i
		},
	) {
		return true
	}

	// check vertical
	if m.checkRows(
		func(i int) (int, int) {
			return m.cursorY - m.streakToWin + 1 + i, m.cursorX
		},
	) {
		return true
	}

	// check diagonal left to right
	if m.checkRows(
		func(i int) (int, int) {
			return m.cursorY - m.streakToWin + 1 + i, m.cursorX - m.streakToWin + 1 + i
		},
	) {
		return true
	}

	// check diagonal right to left
	if m.checkRows(
		func(i int) (int, int) {
			return m.cursorY + m.streakToWin - 1 - i, m.cursorX - m.streakToWin + 1 + i
		},
	) {
		return true
	}

	return false
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.gameOver {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "down":
			if m.cursorY < len(m.board)-1 {
				m.cursorY++
			}
		case "left":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right":
			if m.cursorX < len(m.board[0])-1 {
				m.cursorX++
			}
		case "d":
			if m.debug {
				m.debug = false
			} else {
				m.debug = true
			}
		case "enter", " ":
			cell := &m.board[m.cursorY][m.cursorX]

			if cell.mark == "" {
				cell.mark = m.players[m.currentPlayerId].mark
				m.nextTurn(m.cursorY, m.cursorX)
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var result strings.Builder

	result.WriteString("\n")
	result.WriteString(fmt.Sprintf("%s\n\n\n", m.message))

	for y, row := range m.board {
		for i := range row {
			result.WriteString("┼───")
			if len(row)-1 == i {
				result.WriteString("┼")
			}
		}
		result.WriteString("\n")

		for x, cell := range row {
			if x == m.cursorX && y == m.cursorY {
				if cell.mark != "" {
					result.WriteString(fmt.Sprintf("│█%s█", cell.mark))
				} else {
					result.WriteString("│███")
				}
			} else {
				if cell.mark != "" {
					result.WriteString(fmt.Sprintf("│ %s ", cell.mark))
				} else {
					result.WriteString(fmt.Sprintf("│   "))
				}
			}
			if len(row)-1 == x {
				result.WriteString("│")
			}
		}
		result.WriteString("\n")
	}

	for i := range m.board[0] {
		result.WriteString("┼───")
		if len(m.board[0])-1 == i {
			result.WriteString("┼")
		}
	}

	result.WriteString("\n\n")
	result.WriteString("Press \"enter\" to occupy field.\n")
	result.WriteString("Press \"q\" to quit. \"d\" to debug.\n")
	if m.debug {
		result.WriteString(fmt.Sprintf("Debug: %v\n", m.debug))
		result.WriteString(fmt.Sprintf("Debug message: %v\n", m.debugMessage))
		result.WriteString(fmt.Sprintf("Turns left: %d\n", m.turnsLeft))
		result.WriteString(fmt.Sprintf("Cursor X/Y: %d/%d\n", m.cursorX, m.cursorY))
		result.WriteString(fmt.Sprintf("Player: %v\n", m.players[m.currentPlayerId]))
	}

	result.WriteString("\n\n")

	return result.String()
}
