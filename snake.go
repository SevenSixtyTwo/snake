package main

import terminal "snake/utils"

type snake struct {
	body      []position
	direction direction
}

func newSnake() *snake {
	maxX, maxY := terminal.GetSize()
	pos := position{maxX / 2, maxY / 2}

	return &snake{
		body:      []position{pos},
		direction: start,
	}
}
