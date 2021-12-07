package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func getFuelNeeded(a, b int, fuelType string) int {
	diff := int(math.Abs(float64(a - b)))
	if fuelType == "increasing" {
		return (diff + 1) * diff / 2
	} else {
		return diff
	}
}

func fuelNeededForCrabs(crabPositions []int, position int, fuelType string) int {
	sum := 0
	for _, crabPosition := range crabPositions {
		sum += getFuelNeeded(crabPosition, position, fuelType)
	}

	return sum
}

func align(crabPositions []int, l, r int, fuelType string) {
	if r >= l {
		mid := l + (r-l)/2
		midFuel := fuelNeededForCrabs(crabPositions, mid, fuelType)
		leftFuel := fuelNeededForCrabs(crabPositions, mid-1, fuelType)
		rightFuel := fuelNeededForCrabs(crabPositions, mid+1, fuelType)

		if (mid == 0 || leftFuel > midFuel) && (mid == len(crabPositions)-1 || rightFuel > midFuel) {
			fmt.Println(fuelType, midFuel)
			return
		} else if mid > 0 && leftFuel < midFuel {
			align(crabPositions, l, mid-1, fuelType)
		} else {
			align(crabPositions, mid+1, r, fuelType)
		}
	} else {
		panic("Invalid algorithm implementation!")
	}
}

func alignCrabs(crabPositions []int, l, r int, fuelType string) {
	defer wg.Done()
	align(crabPositions, l, r, fuelType)
}

func main() {
	file, err := os.Open("../input/input7.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	horizontalPositions := strings.Split(strings.TrimSpace(scanner.Text()), ",")

	var positions []int
	for _, value := range horizontalPositions {
		position, _ := strconv.Atoi(value)
		positions = append(positions, position)
	}

	sort.Ints(positions)

	wg.Add(1)
	go alignCrabs(positions, 0, len(positions)-1, "linear")
	wg.Add(1)
	go alignCrabs(positions, 0, len(positions)-1, "increasing")
	wg.Wait()
}
