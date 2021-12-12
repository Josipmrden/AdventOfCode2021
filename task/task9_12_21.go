package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var mu2 sync.Mutex
var RiskSum int = 0
var basinSizes []int = []int{}

type position struct {
	row    int
	column int
}

func isOkayLocation(board [][]int, row, column int) bool {
	if row < 0 || column < 0 || row >= len(board) || column >= len(board[0]) {
		return false
	}

	return true
}

func isPartOfBasin(board [][]int, row, column int) bool {
	return board[row][column] != 9
}

func isVisited(board [][]int, row, column int, visited map[position]bool) bool {
	p := position{row, column}
	_, contains := visited[p]
	return contains
}

func countBasinSize(board [][]int, row, column, pRow, pColumn int, visited map[position]bool) int {
	if !isOkayLocation(board, row, column) {
		return 0
	}

	if board[row][column] < board[pRow][pColumn] {
		return 0
	}

	if !isPartOfBasin(board, row, column) {
		return 0
	}

	if isVisited(board, row, column, visited) {
		return 0
	}

	p := position{row, column}
	visited[p] = true

	counted := 1
	counted += countBasinSize(board, row-1, column, row, column, visited)
	counted += countBasinSize(board, row+1, column, row, column, visited)
	counted += countBasinSize(board, row, column-1, row, column, visited)
	counted += countBasinSize(board, row, column+1, row, column, visited)

	return counted
}

func findBasin(board [][]int, row, column int) {
	defer wg.Done()
	visitedMap := make(map[position]bool)
	basinLength := countBasinSize(board, row, column, row, column, visitedMap)

	mu2.Lock()
	basinSizes = append(basinSizes, basinLength)
	mu2.Unlock()
}

func checkRiskPoint(board [][]int, row, column int) {
	defer wg.Done()
	lowerWanted := 4
	lowerCounted := 0
	if row == 0 || row == len(board)-1 {
		lowerWanted--
	}
	if column == 0 || column == len(board[0])-1 {
		lowerWanted--
	}

	if row != 0 && board[row-1][column] > board[row][column] {
		lowerCounted++
	}
	if row != len(board)-1 && board[row+1][column] > board[row][column] {
		lowerCounted++
	}
	if column != 0 && board[row][column-1] > board[row][column] {
		lowerCounted++
	}

	if column != len(board[0])-1 && board[row][column+1] > board[row][column] {
		lowerCounted++
	}

	if lowerCounted > lowerWanted {
		panic("Invalid implementation!")
	}
	if lowerCounted == lowerWanted {
		wg.Add(1)
		go findBasin(board, row, column)

		mu.Lock()
		RiskSum += (board[row][column] + 1)
		mu.Unlock()
	}
}

func main() {
	file, err := os.Open("../input/input9.txt")
	if err != nil {
		panic(err)
	}

	rows := 0
	columns := 0
	scanner := bufio.NewScanner(file)
	var textBoard []string
	for scanner.Scan() {
		text := scanner.Text()
		if columns == 0 {
			columns = len(text)
		}
		rows++
		textBoard = append(textBoard, text)
	}

	board := make([][]int, rows)
	for i, _ := range textBoard {
		board[i] = make([]int, columns)
		for j, val := range textBoard[i] {
			intVal, _ := strconv.Atoi(string(val))
			board[i][j] = intVal
		}
	}

	for i, _ := range board {
		for j, _ := range board[i] {
			wg.Add(1)
			go checkRiskPoint(board, i, j)
		}
	}
	wg.Wait()

	fmt.Println(RiskSum)

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}
