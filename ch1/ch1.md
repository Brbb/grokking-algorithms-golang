# Introduction to Algorithms
> An algorithm is a set of instructions for accomplishing a task.


## Binary Search
- Its input is a sorted list of elements
- If the element is in the list, returns it, otherwise returns null
- With binary search, you guess the middle number and eliminate half of the remaining numbers every time
    - Think about a number guessing game, from 1 to 100, where you go:
    - 50!
    - too high
    - 25!
    - too high
    - 13!
    - too low
    - 19!
    - that's it!

- Running time: `O(log n)`  
- For any list of _n_, binary search will take _log2 n_
- _log2 240.000 = 18_, because _2^18 ± 240k_`

- Log are the flip of exponential. In the statement _log10 100_ is like asking: "How many 10s do we multiply together to the 100?".

> go test -v binary_search_test.go

## Exercises
1.1
> It would take _log2 128 = 7_

1.2
> It would take _log2 256 =

- How long would take to do a *simple search* for 1 billion elements if each takes 1ms per operation?
> 1.000.000.000ms / 1000 = 1.000.000s
> 1.000.000s / 60 = 16.666m
> 16.666m / 60 = 277h
> 277 / 24 +- 11d
> 11 days!

- And binary search?
> _log2 1.000.000.000 = 2 ^ 30
> 30ms!

## Big O Notation
- The worst case scenario on how many operations its needed to run an algorithm
- Some examples
    - _O(log n)_, *log time*, binary search
    - _O(n)_, *linear time*, simple search
    - _O(n * log n)_, quick sort
    - _O(n^2)_, selection sort
    - _O(n!)_, *factorial time*, traveling salesperson 

## Exercises
1.3
> O(log n)

1.4
> O(n)


1.5
> O(n)

1.6
> O(n)

## Bonus: Binary Search Trees