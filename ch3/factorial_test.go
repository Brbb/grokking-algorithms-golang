package main

import (
	"fmt"
	"testing"
)

func factorial(x int) int {
	if x == 1 {
		return 1
	}

	fmt.Printf("factorial(%d)\n", x-1)
	return x * factorial(x-1)
}

// Golang has no optimization for tail-recursive functions
// https://www.ardanlabs.com/blog/2013/09/recursion-and-tail-calls-in-go_26.html
func tailFactorial(x, acc int) int {
	if x == 1 {
		return acc
	}

	// acc accumulates the result
	fmt.Printf("tailFactorial(%d, %d)\n", x-1, x*acc)
	return tailFactorial(x-1, x*acc)
}

func Test_TailFactorial(t *testing.T) {
	tt := []struct {
		input  int
		output int
	}{
		{
			input:  5,
			output: 120,
		},
	}

	for _, test := range tt {
		t.Run("", func(t *testing.T) {

			result := tailFactorial(test.input, 1)

			if result != test.output {
				t.Errorf("expected %d, got %d", test.output, result)
			}
		})
	}
}

func Test_Factorial(t *testing.T) {
	tt := []struct {
		input  int
		output int
	}{
		{
			input:  5,
			output: 120,
		},
	}

	for _, test := range tt {
		t.Run("", func(t *testing.T) {

			result := factorial(test.input)

			if result != test.output {
				t.Errorf("expected %d, got %d", test.output, result)
			}
		})
	}
}
