package main

import (
	"math/rand"
	"time"

	"github.com/vehsamrak/tictac/internal/client"
)

const (
	height = 15
	width  = 15
)

func main() {
	rand.Seed(time.Now().UnixNano())
	client.NewClient().Run(
		height,
		width,
	)
}
