package advent

import (
	"bufio"
	"log"
	"os"
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
