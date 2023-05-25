package main

import "github.com/vehsamrak/tictac/internal/client"

const (
	height = 10
	width  = 10
)

func main() {
	client.NewClient().Run(height, width)
}
