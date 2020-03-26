# Quicksort

## Divide and Conquer
- Not a algorithm per se, but a tool ir your toolbox to help you figure out problems, it's a way to think about a problem
- Relies heavily in recursion
- _With every recursive call, you have to reduce your problem_
- **Main two steps**:
    - Figure out the base case. This should be the simplest case
    - Divide or decrease your problem until it becomes the base case

## Exercises
4.1
> See sum_test.go

4.2
> See counter_test.go

4.3
> See sum_test.go

4.4
> Base case: the array is just one element, if the element is the guess returns it, otherwise the guess is not in the array
> Recursive case: verify which half of the array contains the guess, and call the function again

## Quicksort
- Define the base case: arrays that have zero or on element, they are sorted because well, or the array is empty or have just one element
- Pick a pivot, an random element in the list and pops it out of the array
- Using the pivot, divide the array in two parts: elements smaller than the pivot and elements bigger than the pivot
- Calls the function itself in the smaller array, appends the pivot, calls the function with the bigger array
- Runtime is _O(n log n)_ in the best case (also, average case) if the pivot is randomly picked, takes _O(n^2)_ in the worst case if pivot is fixed (first or last element, for example)

## Big O notation revised
- Big O notation also have a _constant time_, the time that each operation takes, written as _O(c * [n|n^2|...])_
- The constant time is ignore if algorithms have very distinct run times, like _O(n)_ and _O(n^2)_
- But it's relevant in some cases, like:
    - print_items() loops through a list and prints the element
    - print_items2() loops through a list, prints the element and sleeps for 1 second
    - In these examples, both algorithms are _O(n)_, but the second will take longer because its _constant time_ is relevant
- _Quicksort_ and _merge sort_ aer both _O(n log n)_, but _Quicksort_ has a smaller constant time

## Exercises
4.5
> _O(n)_

4.6
> _O(n)_

4.7
> _O(1)_

4.8
> _O(n^2)_