package main

import (
	"fmt"
	"os"
	"testing"
)

var testdata = []int{
	199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
}

func TestCountIncrease(t *testing.T) {
	answer := 7
	testName := fmt.Sprintf("%+v\n", testdata)
	t.Run(testName, func(t *testing.T) {
		output := countIncrease(testdata)
		if output != answer {
			t.Errorf("got %d, answer %d", output, answer)
		}
	})
}

func TestSlidingWindow(t *testing.T) {
	answer := 5
	testName := fmt.Sprintf("%+v\n", testdata)
	t.Run(testName, func(t *testing.T) {
		output := countSlidingWindow(testdata)
		if output != answer {
			t.Errorf("got %d, answer %d", output, answer)
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