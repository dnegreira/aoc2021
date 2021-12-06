package main

import (
	advent "advent/reuse"
	"fmt"
	"log"
	"path/filepath"
)

func main() {

	inputFile, err := filepath.Abs("input")
	if err != nil {
		log.Fatal(err)
	}
	//Load input from file and convert to int.
	data := advent.InputToInt(advent.LoadInput(inputFile))

	totalIncreases := countIncrease(data)
	fmt.Println("Total: ", totalIncreases)
	slidingWindow := countSlidingWindow(data)
	fmt.Println("Total sliding window changes: ", slidingWindow)
}

func countIncrease(data []int) int {
	var total int
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			total++
		}
	}
	return total
}

func countSlidingWindow(data []int) int {
	var total int
	var last int
	max := len(data) - 3
	for i := 0; i < max; i++ {
		current := data[i] + data[i+1] + data[i+2]
		if current > last {
			total++
		}
		last = current
	}
	return total
}
