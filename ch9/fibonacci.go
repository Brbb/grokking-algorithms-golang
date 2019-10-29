package main

import "fmt"

// https://www.youtube.com/watch?v=vYquumk4nWw

// O(n)
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 1

	for i := 2; i <= len(arr)-1; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[len(arr)-1]
}

// O(n^2)
func fibRecursive(n int) (result int) {
	if n == 1 || n == 2 {
		result = 1
	} else {
		return fib(n-1) + fib(n-2)
	}

	return result

}

func main() {
	n := 1000000000
	fmt.Println(fib(n)) // time 7.68s user 4.50s system 143% cpu 8.481 total
	// fmt.Println(fibRecursive(n)) // time -> don't know, didn't let if finish heheh
}
