package main

import (
	"fmt"
	"testing"
)

func sum(arr []int) (r int) {

	// base case
	if len(arr) == 0 {
		fmt.Println("base case!")
		return 0
	}

	// recursive case
	x, arr := arr[len(arr)-1], arr[:len(arr)-1]
	fmt.Println(x, arr)

	return x + sum(arr)
}

func Test_Sum(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "sum #1",
			input:  []int{3, 4, 5},
			output: 12,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := sum(tC.input)

			if r != tC.output {
				t.Errorf("expected %d, got %d", tC.output, r)
			}
		})
	}
}
