package client

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type cell struct {
}

type model struct {
	board   [][]cell
	cursorX int
	cursorY int
}

func NewModel(height, width int) model {
	board := make([][]cell, height)
	for i := range board {
		board[i] = make([]cell, width)
	}

	return model{
		board: board,
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
			if m.cursorY < len(m.board)-1 {
				m.cursorY++
			}
		case "down":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "left":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right":
			if m.cursorX < len(m.board[0])-1 {
				m.cursorX++
			}
		case "enter":
		}
	}

	return m, nil
}

func (m model) View() string {
	var result strings.Builder

	result.WriteString("\n\n")

	for _, row := range m.board {
		for i := range row {
			result.WriteString("┼───")
			if len(row)-1 == i {
				result.WriteString("┼")
			}
		}
		result.WriteString("\n")

		for i := range row {
			result.WriteString("│   ")
			if len(row)-1 == i {
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

	result.WriteString(fmt.Sprintf("\n\nCursor X/Y: %d/%d\n", m.cursorX, m.cursorY))
	result.WriteString("Press \"q\" to quit.\n\n")

	return result.String()
}
