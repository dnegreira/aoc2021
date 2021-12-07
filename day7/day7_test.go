package main

import (
	"fmt"
	"os"
	"testing"
)



var testdata = []int{
	16,1,2,0,4,2,7,1,2,14,
}

func TestCountFuel(t *testing.T) {
	answer := 37
	testName := fmt.Sprintf("%+v\n", testdata)
	t.Run(testName, func(t *testing.T) {
		output := calcMin(countFuel(testdata, 0, 16, false))
		if output != answer {
			t.Errorf("got %d, answer %d", output, answer)
		}
	})
}

func TestSlidingWindow(t *testing.T) {
	answer := 168
	testname := fmt.Sprintf("%+v\n", testdata)
	t.Run(testname, func(t *testing.T) {
		output := calcMin(countFuel(testdata, 0, 16, true))
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