package client

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Run(height, width, streakToWin int) error {
	p := tea.NewProgram(NewModel(height, width, streakToWin))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occurred: %v", err)
		os.Exit(1)
	}

	return nil
}
