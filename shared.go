package main

import (
	terminal "snake/utils"
	"time"

	"golang.org/x/exp/rand"
)

type position [2]int

type direction int

const velocityChange = time.Millisecond * 30

const (
	north direction = iota
	east
	south
	west
	start
)

// true if positions are same
func positionsAreSame(a, b position) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func randomPosition() position {
	width, height := terminal.GetSize()
	x := rand.Intn(width) + 1
	y := rand.Intn(height) + 1

	return [2]int{x, y}
}
