package main

import "testing"

func counter(arr []int) (r int) {

	// base case
	if len(arr) == 0 {
		return 0
	}

	_, arr = arr[len(arr)-1], arr[:len(arr)-1]
	r++

	return r + counter(arr)
}

func Test_Counter(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "test #1",
			input:  []int{1, 2, 3},
			output: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := counter(tC.input)

			if r != tC.output {
				t.Errorf("expected %d, got %d", tC.output, r)
			}
		})
	}
}
