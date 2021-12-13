package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type foldInstruction struct {
	foldAxis rune
	foldLine int
}

type boardObject struct {
	width int
	height int
	board map[coordinate]bool
}

func printBoard(b boardObject) {

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			_, ok := b.board[coordinate{j, i}]
			if ok {
				fmt.Print("⬛")
			} else {
				fmt.Print("⬜")
			}
		}
		fmt.Println()
	}
}

func fold(b boardObject, instruction foldInstruction) boardObject {
	newBoard := make(map[coordinate]bool)

	var x int
	var y int

	for k, _ := range b.board {
		if instruction.foldAxis == 'x' {
			if k.x > instruction.foldLine {
				x = b.width - k.x - 1
			} else {
				x = k.x
			}
			y = k.y
		} else if instruction.foldAxis == 'y' {
			if k.y > instruction.foldLine {
				y = b.height - k.y - 1
			} else {
				y = k.y
			}
			x = k.x
		}

		c := coordinate{x, y}
		newBoard[c] = true
	}

	newWidth := b.width
	newHeight := b.height
	if instruction.foldAxis == 'x' {
		newWidth = newWidth - instruction.foldLine - 1
	}
	if instruction.foldAxis == 'y' {
		newHeight = newHeight - instruction.foldLine - 1
	}
	return boardObject{newWidth, newHeight, newBoard}
}

func main() {
	file, err := os.Open("../input/input13.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	mode := "coordinates"
	board := make(map[coordinate]bool)
	instructionsByStep := make([]foldInstruction, 0)
	height := 0
	width := 0

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			mode = "fold"
			continue
		}

		if mode == "coordinates" {
			coordinateSplit := strings.Split(strings.TrimSpace(text), ",")
			x, _ := strconv.Atoi(coordinateSplit[0])
			y, _ := strconv.Atoi(coordinateSplit[1])

			if x > width {
				width = x
			}
			if y > height {
				height = y
			}

			c := coordinate{x, y}

			board[c] = true
		} else if mode == "fold" {
			instructions := strings.Split(strings.Split(text, " ")[2], "=")
			foldAxis := instructions[0]
			foldLine, _ := strconv.Atoi(instructions[1])

			instruction := foldInstruction{rune(foldAxis[0]), foldLine}
			instructionsByStep = append(instructionsByStep, instruction)
		}
	}

	bo := boardObject{width + 1, height + 1, board}

	for _, instruction := range instructionsByStep {
		bo = fold(bo, instruction)
		fmt.Println(len(bo.board))
	}
	printBoard(bo)
}
