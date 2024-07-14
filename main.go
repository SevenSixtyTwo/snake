package main

import (
	terminal "snake/utils"
)

func main() {
	game := newGame()
	game.beforeGame()

	for {
		maxX, maxY := terminal.GetSize()

		// calculate new head position
		newHeadPos := game.snake.body[0]

		switch game.snake.direction {
		case north:
			newHeadPos[1]--
		case east:
			newHeadPos[0]++
		case south:
			newHeadPos[1]++
		case west:
			newHeadPos[0]--
		}

		// if you hit the wall, game over
		hitWall := newHeadPos[0] < 1 || newHeadPos[1] < 1 || newHeadPos[0] > maxX || newHeadPos[1] > maxY
		if hitWall {
			game.over()
		}

		for _, snake_pos := range game.snake.body {
			for _, wall_pos := range game.walls {
				if positionsAreSame(snake_pos, wall_pos) {
					game.over()
				}
			}
		}

		// if you run into yourself, game over
		for _, pos := range game.snake.body {
			if positionsAreSame(newHeadPos, pos) {
				game.over()
			}
		}

		// add the new head to the body
		game.snake.body = append([]position{newHeadPos}, game.snake.body...)

		ateFood := positionsAreSame(game.food, newHeadPos)
		if ateFood {
			game.score++
			game.placeNewFood()
		} else {
			game.snake.body = game.snake.body[:len(game.snake.body)-1]
		}

		game.draw()
	}
}
