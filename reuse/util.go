package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LoadInput(filepath string) []string {
	// Loads all the input into memory
	// each day should then process it and
	// convert it as it needs to convert.
	var data []string
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
	log.Fatal(err)
	}
	return data
}

func InputToInt(data []string) []int {
	///Lots of times the input need
	///to be converted to Int.
	///Writing this generic function
	///To always do it the same way.
	var intOutput []int
	for _, v := range data {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Error converting to int: ", err)
		}
		intOutput = append(intOutput, i)
	}
	return intOutput
}