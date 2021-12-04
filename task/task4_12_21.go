package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rowColumn struct {
	row int
	column int
}

type board struct {
	numberReceived int
	numberSum int

	rowWinCount []int
	columnWinCount []int

	stoppedPlaying bool
	
	numberMap map[int]rowColumn
}
func newBoard(rows int, columns int, numbers []int) *board {
	if len(numbers) != rows * columns {
		panic("Number length not correct!")
	}

	b := board{}
	b.numberMap = make(map[int]rowColumn)
	b.rowWinCount = make([]int, rows)
	b.columnWinCount = make([]int, columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			currentNumber := numbers[rows * i + j]
			b.numberMap[currentNumber] = rowColumn{row: i, column: j}
			b.numberSum += currentNumber
		}
	}

	return &b
}
func (b *board) newNumberDrawn(number int) {
	if b.stoppedPlaying {
		return
	}

	b.numberReceived = number
	
	if _, ok := b.numberMap[b.numberReceived]; ok {

		rc := b.numberMap[number]

		b.rowWinCount[rc.row]++
		b.columnWinCount[rc.column]++
		b.numberSum -= b.numberReceived

		if b.rowWinCount[rc.row] == len(b.rowWinCount) || b.columnWinCount[rc.column] == len(b.columnWinCount) {
			b.hasWinningCombination()
		}
	}
}

func (b *board) hasWinningCombination() {
	fmt.Println(b.numberSum * b.numberReceived)
	b.stoppedPlaying = true
	// os.Exit(0)
}

func main() {
	file, err := os.Open("../input/input4.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	var bingoNumbers []int
	var boardNums []int
	var boards []board
	rows := 0
	columns := 0

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			if len(boardNums) != 0 {
				board := newBoard(rows, columns, boardNums)
				boards = append(boards, *board)
			}

			rows = 0
			columns = 0
			boardNums = boardNums[:0]
		} else if i == 0 {
			numberStrings := strings.Split(text, ",")
			for _, val := range numberStrings {
				numberInt, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}

				bingoNumbers = append(bingoNumbers, numberInt)
			}
			i++
		} else {
			rowNumStrings := strings.Fields(text)
			for _, val := range rowNumStrings {
				numberInt, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}

				boardNums = append(boardNums, numberInt)
			}
			rows += 1
			columns = len(rowNumStrings)
		}
	}

	board := newBoard(rows, columns, boardNums)
	boards = append(boards, *board)

	for _, num := range bingoNumbers {
		for brd_idx, _ := range boards {
			brd := &boards[brd_idx]
			brd.newNumberDrawn(num)
		}
	}
}