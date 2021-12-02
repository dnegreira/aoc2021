package main

import (
	advent "advent/reuse"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

func main() {
	var rawData []string

	inputFile, err := filepath.Abs("input")
	if err != nil {
		log.Fatal(err)
	}
	rawData = advent.LoadInput(inputFile)

	// since input is made of int
	// we need to convert from string
	// to int
	var data []int
	for _, v := range rawData {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Error converting to int: ", err)
		}
		data = append(data, i)
	}

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
