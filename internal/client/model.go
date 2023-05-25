package client

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type cell struct {
}

type model struct {
	board [][]cell
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
		case "down":
		case "enter":
		}
	}

	return m, nil
}

func (m model) View() string {
	result := "\n\n"

	for _, row := range m.board {
		for i := range row {
			result += fmt.Sprintf("┼───")
			if len(row)-1 == i {
				result += fmt.Sprintf("┼")
			}
		}
		result += fmt.Sprintf("\n")

		for i := range row {
			result += fmt.Sprintf("│   ")
			if len(row)-1 == i {
				result += fmt.Sprintf("│")
			}
		}
		result += fmt.Sprintf("\n")
	}

	for i := range m.board[0] {
		result += fmt.Sprintf("┼───")
		if len(m.board[0])-1 == i {
			result += fmt.Sprintf("┼")
		}
	}

	result += "\n\nPress \"q\" to quit.\n\n"

	return result
}
