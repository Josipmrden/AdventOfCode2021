package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type position struct {
	row    int
	column int
}

var FlashSum int = 0

func printMatrix(mx [][]int) {
	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i])
	}
}

func isValidLocation(octopuses [][]int, r, c int) bool {
	if r >= 0 && r <= len(octopuses)-1 && c >= 0 && c <= len(octopuses[0])-1 {
		return true
	}

	return false
}

func flash(octopuses [][]int, r, c int, flashed map[position]bool) [][]int {
	p := position{r, c}
	flashed[p] = true
	FlashSum++

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if isValidLocation(octopuses, r+i, c+j) {
				octopuses[r+i][c+j]++
				p := position{r+i, j+c}
				_, didFlash := flashed[p]
				if !didFlash && octopuses[r+i][c+j] > 9 {
					octopuses = flash(octopuses, r+i, c+j, flashed)
				}
			}
		}
	}

	return octopuses
}

func simulateStep(octopuses [][]int) ([][]int, map[position]bool) {
	flashed := make(map[position]bool)
	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[0]); j++ {
			octopuses[i][j]++
			if octopuses[i][j] > 9 {
				
				p := position{i, j}
				_, didFlash := flashed[p]

				if !didFlash {
					octopuses = flash(octopuses, i, j, flashed)
				}
			}
		}
	}

	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[0]); j++ {
			if octopuses[i][j] > 9 {
				octopuses[i][j] = 0
			}
		}
	}

	return octopuses, flashed
}

func main() {
	file, err := os.Open("../input/input11.txt")
	if err != nil {
		panic(err)
	}

	rows := 0
	columns := 0
	scanner := bufio.NewScanner(file)
	var textOctopuses []string
	for scanner.Scan() {
		text := scanner.Text()
		if columns == 0 {
			columns = len(text)
		}
		rows++
		textOctopuses = append(textOctopuses, text)
	}

	octopuses := make([][]int, rows)
	for i, _ := range textOctopuses {
		octopuses[i] = make([]int, columns)
		for j, val := range textOctopuses[i] {
			intVal, _ := strconv.Atoi(string(val))
			octopuses[i][j] = intVal
		}
	}

	for i := 1; true; i++ {
		octopuses, flashed := simulateStep(octopuses)
		if len(flashed) == len(octopuses) * len(octopuses[0]) {
			println("Stop", i)
			break
		}
	}

	printMatrix(octopuses)
	fmt.Println(FlashSum)
}
