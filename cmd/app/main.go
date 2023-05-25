package main

import "github.com/vehsamrak/tictac/internal/client"

const (
	height = 3
	width  = 3
)

func main() {
	client.NewClient().Run(height, width)
}
