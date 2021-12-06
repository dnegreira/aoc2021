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

	inputFile, err := filepath.Abs("input")
	if err != nil {
		log.Fatal(err)
	}
	rawData = advent.LoadInput(inputFile)
	answer1, answer2 := findDepth(rawData)
	fmt.Printf("Answers are for problem1: %d, for problem2: %d\n",answer1, answer2)
}

func findDepth(data []string) (int, int) {
	var distance, depth, aim int
	for _, v := range data {
		order := strings.Fields(v)[0]
		value, err := strconv.Atoi(strings.Fields(v)[1])
		if err != nil {
			log.Fatal("Error converting string to int on: ", err)
		}

		switch order {
		case "forward":
			distance += value
			aim += depth * value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}
	return distance * depth, distance * aim
}