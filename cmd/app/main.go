package main

import (
	"math/rand"
	"time"

	"github.com/vehsamrak/tictac/internal/client"
)

const (
	streakToWin = 5
	height      = 15
	width       = 25
)

func main() {
	rand.Seed(time.Now().UnixNano())
	client.NewClient().Run(
		height,
		width,
		streakToWin,
	)
}
