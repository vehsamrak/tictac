package main

import "github.com/vehsamrak/tictac/internal/client"

const (
	height = 5
	width  = 5
)

func main() {
	client.NewClient().Run(
		height,
		width,
	)
}
