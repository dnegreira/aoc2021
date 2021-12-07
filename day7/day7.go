package main

import (
	advent "advent/reuse"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var rawData []string
	var data []int
	inputFile, err := filepath.Abs("input")
	if err != nil {
		log.Fatal(err)
	}
	rawData = advent.LoadInput(inputFile)
	for _, v := range strings.Split(rawData[0], ",") {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, num)
	}

	min, max := calcMinMaxInput(data)

	regularResult := countFuel(data, min, max, false)
	minRegularResult := calcMin(regularResult)

	multipliedResults := countFuel(data, min, max, true)
	minMultipliedResult := calcMin(multipliedResults)
	fmt.Println("Regular result: ", minRegularResult)
	fmt.Println("Multiplied result: ", minMultipliedResult)
}


func countFuel (data []int, min int, max int, multiplier bool) map[string]int {
	allResults := make(map[string]int)
	// For easy reading :-)
	curNum := min
	for curNum <= max {
		indexNum := strconv.FormatInt(int64(curNum), 10)
		allResults[indexNum] = 0
		var difference int
		for _, nextNum := range data {
			if curNum != nextNum {
				if curNum > nextNum {
					difference = curNum - nextNum
				}
				if curNum < nextNum {
					difference = nextNum - curNum
				}
				if multiplier {
					diffMultiplied := difference
					for i := 0; i < difference; i++ {
						diffMultiplied += i
					}
					indexNum := strconv.FormatInt(int64(curNum), 10)
					allResults[indexNum] += diffMultiplied
				} else {
					indexNum := strconv.FormatInt(int64(curNum), 10)
					allResults[indexNum] += difference
				}
			}
		}
		curNum++
	}
	return allResults
}

func calcMin (allResults map[string]int) int {
	var minResult int
	for i, min := range allResults{
		if minResult == 0 {
			minResult = allResults[i]
		}
		if min < minResult {
			minResult = min
		}
	}
	return minResult
}

func calcMinMaxInput (data []int) (int, int) {
	var min, max int
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}