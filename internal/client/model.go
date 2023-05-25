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
	x    int
	y    int
}

type Player struct {
	mark string
}

func (m *model) nextTurn() {
	if m.currentPlayerId == len(m.players)-1 {
		m.currentPlayerId = 0
	} else {
		m.currentPlayerId++
	}

	m.message = fmt.Sprintf("Current turn: %s", m.players[m.currentPlayerId].mark)
}

type model struct {
	debugMessage    string // string with any message to render in debug mode
	message         string
	board           [][]cell
	players         []Player
	debug           bool
	currentPlayerId int // id is players index
	cursorX         int
	cursorY         int
}

func NewModel(height, width int) model {
	board := make([][]cell, height)
	for y := 0; y < height; y++ {
		board[y] = make([]cell, width)
		for x := 0; x < width; x++ {
			board[y][x] = cell{
				x: x,
				y: y,
			}
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
			// mark: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colors[i].Hex())).Render(mark),
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
		message: fmt.Sprintf(
			"New game started! Goal is to occupy 5 in a row. Current turn: %s",
			players[currentPlayerId].mark,
		),
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
			if m.debug {
				m.debug = false
			} else {
				m.debug = true
			}
		case "enter", " ":
			cell := &m.board[m.cursorY][m.cursorX]

			if cell.mark == "" {
				cell.mark = m.players[m.currentPlayerId].mark
				m.nextTurn()
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
				if m.debug {
					result.WriteString(fmt.Sprintf("│%d %d", cell.x, cell.y))
				} else {
					if cell.mark != "" {
						result.WriteString(fmt.Sprintf("│█%s█", cell.mark))
					} else {
						result.WriteString("│███")
					}
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
		result.WriteString(fmt.Sprintf("Cursor X/Y: %d/%d\n", m.cursorX, m.cursorY))
		result.WriteString(fmt.Sprintf("Player: %v\n", m.players[m.currentPlayerId]))
	}

	result.WriteString("\n\n")

	return result.String()
}
