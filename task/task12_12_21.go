package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var network map[string][]string = make(map[string][]string)
var startPoint = "start"
var endPoint = "end"

func isSmallCave(s string) bool {
	if isStartOrEnd(s) {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func isStartOrEnd(s string) bool {
	return startPoint == s || endPoint == s
}

func isBigCave(s string) bool {
	return !isStartOrEnd(s) && !isSmallCave(s)
}

func newVisitedMap(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func hasAlreadyVisitedSmallCaveTwice(m map[string]int) bool{
	for k, v := range m {
		if isSmallCave(k) {
			if v >= 2 {
				return true
			}
		}
	}

	return false
}

func canVisit(node string, m map[string]int, bonus bool) bool {
	if isBigCave(node) {
		return true
	}

	val, ok := m[node]
	if !ok {
		return true
	}

	if isStartOrEnd(node) && val >= 1 {
		return false
	}

	if isSmallCave(node) {
		if bonus {
			if hasAlreadyVisitedSmallCaveTwice(m) {
				if val >= 1 {
					return false
				}
			} else {
				if val >= 2 {
					return false
				}
			}
		} else {
			if val >= 1 {
				return false
			}
		}
	}

	return true
}

func visitNode(node string, visited map[string]int) {
	_, ok := visited[node]
	if !ok {
		visited[node] = 0
	}
	visited[node]++
}

func countPaths(node string, visited map[string]int, bonus bool) int {
	if node == endPoint {
		return 1
	}

	visitNode(node, visited)

	pathsCounted := 0
	for _, neighbor := range network[node] {
		if canVisit(neighbor, visited, bonus) {
			newVisited := newVisitedMap(visited)
			pathsCounted += countPaths(neighbor, newVisited, bonus)
		}
	}

	return pathsCounted
}

func main() {
	file, err := os.Open("../input/input12.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), "-")
		start := edge[0]
		end := edge[1]

		_, contained := network[start]
		if !contained {
			network[start] = make([]string, 0)
		}
		_, contained = network[end]
		if !contained {
			network[end] = make([]string, 0)
		}

		network[start] = append(network[start], end)
		network[end] = append(network[end], start)
	}

	fmt.Println(network)
	visited := make(map[string]int)
	noPaths := countPaths(startPoint, visited, false)
	fmt.Println(noPaths)

	visited2 := make(map[string]int)
	noPaths2 := countPaths(startPoint, visited2, true)
	fmt.Println(noPaths2)
}
