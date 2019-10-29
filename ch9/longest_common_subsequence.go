package main

import "fmt"

func main() {

	s1 := "fish"
	s2 := "fosh"

	grid := make([][]int, len(s2))
	for i := 0; i < len(s1); i++ {
		grid[i] = make([]int, len(s2))
	}

	maxValue := 0
	for i, s1_char := range s1 {
		for j, s2_char := range s2 {
			if s1_char == s2_char {
				maxValue++
				if j-1 >= 0 {
					grid[i][j] = maxValue + grid[i][j-1]
				} else {
					grid[i][j] = maxValue
				}
			} else {
				grid[i][j] = maxValue
			}
		}
	}

	fmt.Printf("Longest common subsequence size between %q and %q is %d!\n", s1, s2, maxValue)
}
