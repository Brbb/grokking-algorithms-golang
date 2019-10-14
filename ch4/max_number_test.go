package main

import (
	"fmt"
	"testing"
)

// maxNumber compares the first two elements of a slice
func maxNumber(arr []int) (r int) {
	fmt.Println("called")

	// base case
	if len(arr) == 2 {
		fmt.Println("base case")
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}

	// recursive case
	subMax := maxNumber(arr[1:])

	if arr[0] > subMax {
		fmt.Println("arr[0] > subMax")
		return arr[0]
	} else {
		fmt.Println("arr[0] < subMax")
		return subMax
	}
}

// maxNumber tests the last two elements and remove the smaller one,
// until the array is just two elements, compraring them as well in the base case
func maxNumber2(arr []int) int {
	// base case
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}

	// recursive case
	if arr[len(arr)-1] >= arr[len(arr)-2] {
		// remove arr[-2]
		arr = append(arr[:len(arr)-2], arr[len(arr)-2+1:]...)
		return maxNumber2(arr)
	}

	if arr[len(arr)-1] <= arr[len(arr)-2] {
		// remove arr[-1], pop
		_, arr = arr[len(arr)-1], arr[:len(arr)-1]
		return maxNumber2(arr)
	}

	return 0
}

func Test_MaxNumber(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "test #1",
			input:  []int{1, 100, 9, 2, 0, 23, 59},
			output: 100,
		},
		{
			desc:   "test #2",
			input:  []int{1, 9, 2, 0, 23, 59, 100},
			output: 100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := maxNumber(tC.input)

			if r != tC.output {
				t.Errorf("expected %d, got %d", tC.output, r)
			}
		})
	}
}

func Test_MaxNumber2(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "test #1",
			input:  []int{1, 100, 9, 2, 0, 23, 59},
			output: 100,
		},
		{
			desc:   "test #2",
			input:  []int{1, 9, 2, 0, 23, 59, 100},
			output: 100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := maxNumber2(tC.input)

			if r != tC.output {
				t.Errorf("expected %d, got %d", tC.output, r)
			}
		})
	}
}
