package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu1 sync.Mutex
var mu2 sync.Mutex

var Sum1 int = 0
var Sum2 int = 0

type insights struct {
	LettersIn1 string
	LettersIn4 string
	LettersIn7 string
}

func letterSimilarity(s1 string, s2 string) int {
	similarity := 0
	for _, rune1 := range s1 {
		for _, rune2 := range s2 {
			if rune1 == rune2 {
				similarity++
			}
		}
	}

	return similarity
}

func getNumber(encoded string, i insights) int {
	if len(encoded) == 2 {
		return 1
	} else if len(encoded) == 3 {
		return 7
	} else if len(encoded) == 4 {
		return 4
	} else if len(encoded) == 7 {
		return 8
	} else if len(encoded) == 5 {
		if letterSimilarity(encoded, i.LettersIn7) == 3 {
			return 3
		} else if letterSimilarity(encoded, i.LettersIn4) == 3 {
			return 5
		} else {
			return 2
		}
	} else if len(encoded) == 6 {
		if letterSimilarity(encoded, i.LettersIn1) == 1 {
			return 6
		} else if letterSimilarity(encoded, i.LettersIn4) == 4 {
			return 9
		} else {
			return 0
		}
	} else {
		panic("Did not cover case!")
	}
}

func getInsights(sample []string) insights {
	i := insights{}
	for _, number := range sample {
		if len(number) == 2 {
			i.LettersIn1 = number
		} else if len(number) == 3 {
			i.LettersIn7 = number
		} else if len(number) == 4 {
			i.LettersIn4 = number
		}
	}

	return i
}

func decode(encodedSamples []string, encodedOutput []string) {
	defer wg.Done()
	i := getInsights(encodedSamples)

	mu2.Lock()
	Sum2 += 1000*getNumber(encodedOutput[0], i) + 100*getNumber(encodedOutput[1], i) + 10*getNumber(encodedOutput[2], i) + getNumber(encodedOutput[3], i)
	mu2.Unlock()
}

func count1478(encodedOutput []string) {
	defer wg.Done()
	for _, output := range encodedOutput {
		if Is1478(output) {
			mu1.Lock()
			Sum1++
			mu1.Unlock()
		}
	}
}

func Is1478(s string) bool {
	if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("../input/input8.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sampleWithOutput := strings.Split(line, "|")
		sample := strings.TrimSpace(sampleWithOutput[0])
		output := strings.TrimSpace(sampleWithOutput[1])
		encodedSamples := strings.Split(sample, " ")
		encodedOutput := strings.Split(output, " ")

		wg.Add(1)
		go count1478(encodedOutput)
		wg.Add(1)
		go decode(encodedSamples, encodedOutput)
	}

	wg.Wait()

	fmt.Println(Sum1)
	fmt.Println(Sum2)
}
