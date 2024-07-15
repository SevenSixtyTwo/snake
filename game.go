package main

import (
	"log"
	"os"
	"os/signal"
	terminal "snake/utils"
	"strconv"
	"time"

	"github.com/mattn/go-tty"
	"golang.org/x/exp/rand"
)

type game struct {
	score         int
	snake         *snake
	food          position
	speed         time.Duration
	speedModified bool
	walls         []position
}

func newGame() *game {
	rand.Seed(uint64(time.Now().UnixNano()))

	snake := newSnake()

	wallsCount := 20
	var walls []position

	for i := 0; i < wallsCount; i++ {
		newWall := randomPosition()
		for _, pos := range walls {
			for positionsAreSame(newWall, pos) {
				newWall = randomPosition()
			}
		}
		walls = append(walls, newWall)
	}

	food := randomPosition()

	for _, pos := range walls {
		for positionsAreSame(food, pos) {
			food = randomPosition()
		}
	}

	game := &game{
		score:         0,
		snake:         snake,
		food:          food,
		speed:         time.Millisecond * 50,
		speedModified: true,
		walls:         walls,
	}

	go game.listenForKeyPress()

	return game
}

func (g *game) beforeGame() {
	terminal.HideCursor()

	// handle ^C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			g.over()
		}
	}()
}

func (g *game) over() {
	terminal.Clear()
	terminal.ShowCursor()

	terminal.MoveCursor(position{1, 1})
	terminal.Draw("game over. score: " + strconv.Itoa(g.score) + "\n")

	terminal.Render()

	terminal.ReturnEcho()

	os.Exit(0)
}

func (g *game) draw() {
	terminal.Clear()
	maxX, _ := terminal.GetSize()

	status := "score: " + strconv.Itoa(g.score)
	statusXPos := maxX/2 - len(status)/2

	terminal.MoveCursor(position{statusXPos, 0})
	terminal.Draw(status)

	speed := "speed: " + strconv.Itoa(int(g.speed))

	terminal.MoveCursor(position{0, 0})
	terminal.Draw(speed)

	terminal.MoveCursor(g.food)
	terminal.Draw("\033[33m*\033[0m")

	for _, pos := range g.walls {
		terminal.MoveCursor(pos)
		terminal.Draw("\033[31m#\033[0m")
	}

	for i, pos := range g.snake.body {
		terminal.MoveCursor(pos)

		if i == 0 {
			terminal.Draw("\033[32mO\033[0m")
		} else {
			terminal.Draw("\033[32mo\033[0m")
		}
	}

	terminal.Render()
	time.Sleep(g.speed)
}

func (g *game) placeNewFood() {
	for {
		newFoodPosition := randomPosition()

		if positionsAreSame(newFoodPosition, g.food) {
			continue
		}

		for _, pos := range g.walls {
			for positionsAreSame(newFoodPosition, pos) {
				newFoodPosition = randomPosition()
			}
		}

		for _, pos := range g.snake.body {
			if positionsAreSame(newFoodPosition, pos) {
				continue
			}
		}

		g.food = newFoodPosition

		break
	}
}

func (g *game) listenForKeyPress() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		char, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		// UP, DOWN, RIGHT, LEFT = [A, [B, [C, [D
		// we ignore the escape character [
		switch char {
		case 'A':
			if !g.speedModified {
				g.speedModified = true
				g.speed += velocityChange
			}
			g.snake.direction = north
		case 'B':
			if !g.speedModified {
				g.speedModified = true
				g.speed += velocityChange
			}
			g.snake.direction = south
		case 'C':
			if g.speedModified {
				g.speedModified = false
				g.speed -= velocityChange
			}
			g.snake.direction = east
		case 'D':
			if g.speedModified {
				g.speedModified = false
				g.speed -= velocityChange
			}
			g.snake.direction = west
		}
	}
}
