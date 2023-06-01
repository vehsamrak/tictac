package client

import (
	"fmt"
	"strings"

	"github.com/vehsamrak/tictac/internal/minimax"
	"github.com/vehsamrak/tictac/internal/tictac"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

type Player struct {
	mark  string
	isBot bool
}

type model struct {
	messageDebug string // string with any message to render in debug mode
	message      string // game message to render
	board        [][]string
	players      []Player // rotating slice of all players, first is those who's turn it is
	debug        bool
	gameOver     bool
	turnsLeft    int
	cursorX      int
	cursorY      int
	streakToWin  int // marks streak needed to win
}

func NewModel(height, width, streakToWin int) model {
	board := make([][]string, height)
	for y := 0; y < height; y++ {
		board[y] = make([]string, width)
	}

	var usedColors []colorful.Color
	players := []Player{
		{mark: "X"},
		{mark: "O", isBot: true},
	}
	for i, player := range players {
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

		players[i].mark = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(color.Hex())).
			Render(player.mark)
	}

	return model{
		board:       board,
		cursorX:     width / 2,
		cursorY:     height / 2,
		players:     players,
		streakToWin: streakToWin,
		turnsLeft:   height * width,
		message: fmt.Sprintf(
			"New game started! Goal is to occupy %d in a row. Current turn: %s",
			streakToWin,
			players[0].mark,
		),
	}
}

func (m *model) nextTurn(y int, x int) {
	if tictac.CheckGameOver(m.board, y, x, m.streakToWin) {
		m.message = fmt.Sprintf("Game over. Winner is %s!", m.players[0].mark)
		m.gameOver = true
		return
	}

	m.turnsLeft--
	if m.turnsLeft == 0 {
		m.message = fmt.Sprintf("Game over. Draw!")
		m.gameOver = true
		return
	}

	// players rotate
	m.players = append(m.players[1:], m.players[0])

	m.message = fmt.Sprintf("Current turn: %s", m.players[0].mark)

	if m.players[0].isBot {
		var marks []string
		for _, player := range m.players {
			marks = append(marks, player.mark)
		}

		prediction, predictedY, predictedX := minimax.Minimax{}.Minimax(
			minimax.Data{
				Players:       marks,
				CursorY:       m.cursorY,
				CursorX:       m.cursorX,
				MaximizerMark: m.players[0].mark,
				StreakToWin:   m.streakToWin,
			},
			m.board,
			0,
		)

		m.messageDebug = fmt.Sprintf("Prediction: %d", prediction)

		m.placeMark(predictedY, predictedX)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.debug = !m.debug
		case "enter", " ":
			// if m.board[m.cursorY][m.cursorX] == "" {
			// 	m.board[m.cursorY][m.cursorX] = m.players[m.currentPlayerID].mark
			// 	m.nextTurn(m.cursorY, m.cursorX)
			// }

			m.placeMark(m.cursorY, m.cursorX)

			if m.gameOver {
				return m, tea.Quit
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
				if cell != "" {
					result.WriteString(fmt.Sprintf("│█%s█", cell))
				} else {
					result.WriteString("│███")
				}
			} else {
				if cell != "" {
					result.WriteString(fmt.Sprintf("│ %s ", cell))
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
		result.WriteString(fmt.Sprintf("Debug message: %v\n", m.messageDebug))
		result.WriteString(fmt.Sprintf("Turns left: %d\n", m.turnsLeft))
		result.WriteString(fmt.Sprintf("Cursor X/Y: %d/%d\n", m.cursorX, m.cursorY))
		result.WriteString(fmt.Sprintf("Player: %v\n", m.players[0]))
	}

	result.WriteString("\n\n")

	return result.String()
}

func (m *model) placeMark(y int, x int) {
	if m.board[y][x] == "" {
		m.board[y][x] = m.players[0].mark
		m.nextTurn(y, x)
	}
}
