package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var data []int
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		s, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	totalIncreases := countIncrease(data)
	fmt.Println("Total: ", totalIncreases)
	slidingWindow := countSlidingWindow(data)
	fmt.Println("Total sliding window changes: ", slidingWindow)
}


func countIncrease(data []int)  int {
	var total int
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			total++
		}
	}
	return total
}

func countSlidingWindow(data []int)  int {
	var total int
	max := len(data) - 3
	for i := 0; i < max; i++ {
		next := data[i+1] + data[i+2] + data[i+3]
		current := data[i] + data[i+1] + data[i+2]
		if current < next {
			total++
		}
	}
	return total
}