package client

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type cell struct {
	isOccupied bool
	x          int
	y          int
}

type model struct {
	board   [][]cell
	cursorX int
	cursorY int
	debug   bool
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

	return model{
		board:   board,
		cursorX: width / 2,
		cursorY: height / 2,
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
		case "enter":
		}
	}

	return m, nil
}

func (m model) View() string {
	var result strings.Builder

	result.WriteString("\n\n")

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
					result.WriteString("│███")
				}
			} else {
				result.WriteString(fmt.Sprintf("│   "))
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

	result.WriteString("\n\nPress \"q\" to quit. \"d\" to debug.\n")
	if m.debug {
		result.WriteString(fmt.Sprintf("Cursor X/Y: %d/%d\n", m.cursorX, m.cursorY))
		result.WriteString(fmt.Sprintf("Debug: %v\n", m.debug))
	}

	result.WriteString("\n\n")

	return result.String()
}
