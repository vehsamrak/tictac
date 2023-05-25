package main

import "github.com/vehsamrak/tictac/internal/client"

const (
	height = 15
	width  = 15
)

func main() {
	client.NewClient().Run(
		height,
		width,
	)
}
