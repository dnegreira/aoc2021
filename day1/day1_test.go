package main

import (
	"fmt"
	"testing"
)

var testdata = []int{
	199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
}

func TestCountIncrease(t *testing.T) {
	answer := 7
	testname := fmt.Sprintf("%+v\n", testdata)
	t.Run(testname, func(t *testing.T) {
		output := countIncrease(testdata)
		if output != answer {
			t.Errorf("got %d, answer %d", output, answer)
		}
	})
}

func TestSlidingWindow(t *testing.T) {
	answer := 5
	testname := fmt.Sprintf("%+v\n", testdata)
	t.Run(testname, func(t *testing.T) {
		output := countSlidingWindow(testdata)
		if output != answer {
			t.Errorf("got %d, answer %d", output, answer)
		}
	})
}