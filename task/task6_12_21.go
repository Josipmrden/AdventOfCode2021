package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(array []int) int {  
	result := 0  
	for _, v := range array {  
	 result += v  
	}  
	return result  
}

func calculateFishSpread(fish []int, noDays int) {
	for i := 0; i < noDays; i++ {
		newLanternFish := fish[0]
		for j := 0; j < len(fish) - 1; j++ {
			fish[j] = fish[j+1]
		}

		fish[6] += newLanternFish
		fish[8] = newLanternFish
	}

	fmt.Println("Result :", sum(fish))
}

func main() {
	file, err := os.Open("../input/input6.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fishStrings := strings.Split(strings.TrimSpace(scanner.Text()), ",")

	fish := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fishTimer := range fishStrings {
		fishTimerInt, _ := strconv.Atoi(fishTimer)
		fish[fishTimerInt]++
	}

	calculateFishSpread(fish, 256)
}
