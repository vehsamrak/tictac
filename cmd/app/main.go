package main

import (
	"math/rand"
	"time"

	"github.com/vehsamrak/tictac/internal/client"
)

const (
	streakToWin = 3
	height      = 3
	width       = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	client.NewClient().Run(
		height,
		width,
		streakToWin,
	)
}
