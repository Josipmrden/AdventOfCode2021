package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type line struct {
	start coordinate
	end   coordinate
}

//////////////////// BOARD //////////////////////////

type board struct {
	boardMap               map[coordinate]int
	multipleOccurenceCount int
}

func newBoard() *board {
	board := board{}
	board.boardMap = make(map[coordinate]int)

	return &board
}

func (board *board) printMaxOccurences() {
	fmt.Println(board.multipleOccurenceCount)
}

func (board *board) addCoordinate(c coordinate) {
	if _, found := board.boardMap[c]; !found {
		board.boardMap[c] = 0
	}

	board.boardMap[c] += 1

	occurence, _ := board.boardMap[c]
	if occurence == 2 {
		board.multipleOccurenceCount++
	}
}

func (board *board) addLine(l line) {
	rangeLength := int(math.Max(math.Abs(float64(l.start.x - l.end.x)), math.Abs(float64(l.start.y - l.end.y)))) + 1
	xRange := makeRange(l.start.x, l.end.x, rangeLength)
	yRange := makeRange(l.start.y, l.end.y, rangeLength)

	for idx, _ := range xRange {
		c := coordinate{xRange[idx], yRange[idx]}
		board.addCoordinate(c)
	}
}

////////////////////////// UTILS ////////////////////////////////////

func makeRange(a, b, length int) []int {
	arr := make([]int, length)

	for i := range arr {
		if a == b {
			arr[i] = a
		} else if a > b {
			arr[i] = a - i
		} else {
			arr[i] = a + i
		}
	}

	return arr
}

func parseCoordinate(text string) coordinate {
	trimmedText := strings.TrimSpace(text)
	splittedCoordinate := strings.Split(trimmedText, ",")

	xCoordinate, _ := strconv.Atoi(strings.TrimSpace(splittedCoordinate[0]))
	yCoordinate, _ := strconv.Atoi(strings.TrimSpace(splittedCoordinate[1]))

	coordinate := coordinate{xCoordinate, yCoordinate}

	return coordinate
}

func parseBoard() *board {
	file, err := os.Open("../input/input5.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	board := newBoard()

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		coordinatePairStrings := strings.Split(text, " -> ")

		c1 := parseCoordinate(coordinatePairStrings[0])
		c2 := parseCoordinate(coordinatePairStrings[1])
		l := line{c1, c2}

		board.addLine(l)
	}

	return board
}

///////////////////////////////////////////////////////7

func main() {
	board := parseBoard()

	board.printMaxOccurences()
}
