package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var mu2 sync.Mutex
var IllegalSum int = 0
var LegalScores []int = make([]int, 0)

func awardIllegalPoints(r rune) int {
	if strings.Contains("()", string(r)) {
		return 3
	} else if strings.Contains("[]", string(r)) {
		return 57
	} else if strings.Contains("{}", string(r)) {
		return 1197
	} else if strings.Contains("<>", string(r)) {
		return 25137
	}

	panic("Invalid award implementation!")
}

func awardLegalPoints(r rune) int {
	if strings.Contains("()", string(r)) {
		return 1
	} else if strings.Contains("[]", string(r)) {
		return 2
	} else if strings.Contains("{}", string(r)) {
		return 3
	} else if strings.Contains("<>", string(r)) {
		return 4
	}

	panic("Invalid award implementation!")
}

func isCorrectBrace(r1 rune, r2 rune) bool {
	if r1 == '(' && r2 == ')' {
		return true
	}
	if r1 == '[' && r2 == ']' {
		return true
	}
	if r1 == '{' && r2 == '}' {
		return true
	}
	if r1 == '<' && r2 == '>' {
		return true
	}

	return false
}

func isBeginRune(r1 rune) bool {
	return strings.Contains("([{<", string(r1))
}

func addLegalScore(stack []rune) {
	score := 0

	for len(stack) != 0 {
		n := len(stack) - 1
		topElement := stack[n]
		score = score * 5 + awardLegalPoints(topElement)

		stack = stack[:n]
	}

	mu2.Lock()
	LegalScores = append(LegalScores, score)
	mu2.Unlock()
}

func findIllegal(s string) {
	defer wg.Done()

	var stack []rune
	for _, r := range s {
		if isBeginRune(r) {
			stack = append(stack, r)
		} else {
			n := len(stack) - 1
			topElement := stack[n]
			if !isCorrectBrace(topElement, r) {
				mu.Lock()
				IllegalSum += awardIllegalPoints(r)
				mu.Unlock()

				return
			}
			stack = stack[:n]
		}
	}
	addLegalScore(stack)
}

func main() {
	file, err := os.Open("../input/input10.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		wg.Add(1)
		go findIllegal(text)
	}

	wg.Wait()
	fmt.Println(IllegalSum)
	sort.Ints(LegalScores)
	fmt.Println(LegalScores)
	fmt.Println(len(LegalScores))
	fmt.Println(LegalScores[len(LegalScores) / 2])
}