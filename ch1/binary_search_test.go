package main

import (
	"fmt"
	"testing"
)

func binarySearch(list []int, item int) (pos, steps int) {
	low := 0
	high := len(list) - 1

	for high >= low {
		steps += 1

		mid := (low + high) / 2
		guess := list[mid]

		fmt.Printf("step %d, high %d, low %d, mid %d, item %d, guess %d\n", steps, high, low, mid, item, guess)

		switch {
		case guess == item:
			return mid, steps
		// move our high pointer to the middle, discards right side
		case guess > item:
			high = mid - 1
		// move our low pointer to the middle, discards left side
		case guess < item:
			low = mid + 1
		}
	}

	return -1, steps
}

func makeRange(size int) (r []int) {
	var i int = 1

	for i <= size {
		r = append(r, i)
		i += 1
	}

	return r
}

func Test_BinarySearch(t *testing.T) {
	type in struct {
		list []int
		item int
	}
	type Test struct {
		in    in
		out   int
		steps int
	}

	var tableTest = []Test{
		{
			in:    in{[]int{1, 2, 3, 4, 5}, 5},
			out:   4,
			steps: 3,
		},
		{
			in:    in{[]int{1, 2, 3, 4}, 10},
			out:   -1,
			steps: 3,
		},
		{
			in:    in{makeRange(240000), 240000},
			out:   239999,
			steps: 18,
		},
		{
			in:    in{makeRange(240000), 1},
			out:   0,
			steps: 17,
		},
	}

	for _, test := range tableTest {
		t.Run(fmt.Sprintf("list size %d, item %d", len(test.in.list), test.in.item), func(t *testing.T) {
			pos, steps := binarySearch(test.in.list, test.in.item)
			if pos != test.out {
				t.Errorf("pos: expected %d, got %d", test.out, pos)
			}
			if steps != test.steps {
				t.Errorf("steps: expected %d, got %d", test.steps, steps)
			}
		})
	}
}
