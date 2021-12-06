package main

import (
	"fmt"
	"os"
	"testing"
)

var testdata = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestFindDepth(t *testing.T) {
	answer1 := 150
	answer2 := 900
	testName := fmt.Sprintf("%+v\n", testdata)
	t.Run(testName, func(t *testing.T) {
		output1, output2 := findDepth(testdata)
		if output1 != answer1 {
			t.Errorf("got %d, answer %d", output1, answer1)
		}
		if output2 != answer2 {
			t.Errorf("got %d, answer %d", output2, answer2)
		}
	})
}

func BenchmarkMain1000(b *testing.B) {
	// run the Fib function b.N times
	os.Stdout = nil
	os.Stderr = nil
	for n := 0; n < b.N; n++ {
		main()
	}
}