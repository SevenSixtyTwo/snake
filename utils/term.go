package terminal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

// buffer that I will write to
var screen = bufio.NewWriter(os.Stdout)

// using ANSI escape codes for various tasks
func HideCursor() {
	fmt.Fprint(screen, "\033[?25l")
}

func ShowCursor() {
	fmt.Fprint(screen, "\033[?25h")
}

func MoveCursor(pos [2]int) {
	fmt.Fprintf(screen, "\033[%d;%dH", pos[1], pos[0])
}

func ReturnEcho() {
	cmd := exec.Command("stty", "echo")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	fmt.Fprint(screen, "\033[2J")

	// cmd := exec.Command("clear")
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Run()
}

func Draw(str string) {
	fmt.Fprint(screen, str)
}

func Render() {
	screen.Flush()
}

// top left position is 1:1
// return: width, height
func GetSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	return width, height
}
