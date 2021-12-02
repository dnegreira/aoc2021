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

	fmt.Println("Answer for problem1 is: ", findDepth(rawData))
	fmt.Println("Answer for problem2 is: ", findDepthWithAim(rawData))
}

func findDepth(data []string) int {
	var distance, depth int
	for _, v := range data {
		order := strings.Fields(v)[0]
		value, err := strconv.Atoi(strings.Fields(v)[1])
		if err != nil {
			log.Fatal("Error converting string to int on: ", err)
		}

		switch order {
		case "forward":
			distance += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}
	return distance * depth
}

func findDepthWithAim(data []string) int {
	var distance, depth, aim int
	for _, v := range data {
		order := strings.Fields(v)[0]
		value, err := strconv.Atoi(strings.Fields(v)[1])
		if err != nil {
			log.Fatal("Error converting string to int: ", err)
		}

		switch order {
		case "forward":
			distance += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	return distance * depth
}
