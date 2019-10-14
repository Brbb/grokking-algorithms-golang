# Recursion
- A function that calls itself!
- All recursion functions has two parts:
    - Base case: condition to exit 
    - Recursive case: when the function calls itself

## Stack
- A data structure where you add items on top, _push_, and remove the topmost item, _pop_
- Like a todo list where you can just add things on top and read the last added item
- Why is important? Because of the call stack!

## Call Stack
- A Call Stack is the "computer internally" stack, stored in memory, where it keep tracks of what functions to execute and its data.
- In a recursive function, the stack will keep growing each time the _recursive case_ is executed, pushing the function and its values. When the function hits the _base case_ item by item is popped and computed.
- The only drawback of recursion (+stack) is that we need to take care of how much memory we are consuming to not explode the stack, _stack overflow_

## Exercise
3.1
> A function named greet was called with a variable named "name" with "maggie" value. Inside this function, another function named "greet2" was called with a variable named "name" with "maggie" value.

## Resources
https://www.ardanlabs.com/blog/2013/09/recursion-and-tail-calls-in-go_26.html
