# Greedy Algorithms
- Algorithms to be used with very hard to solve problems
- They don't give you a perfect answer, but a good enough one

- Instead of finding the perfect path or set at every iteration, simply get the better one
- Finding the locally optimal solution in each iteration will lead you to the globally optimal solution

## Exercises
8.1:
> Get the biggest boxes first, put in the trunk, repeat. Not the optimal solution, but can carry the biggest stuff first!

8.2:
> Get the item with most point and lowest distance, add to the list, repeat. Not optimal, but well.

## The set-covering problem
- You have a radio show and want to play in all 50 US states
- You want to minimize the number of stations you play, because they cost money
- Each stations plays in one or more states
- Very hard problem to solve!

1. List all the possible subsets of stations. This is called a _power set_, there are _2^n_ possible stations
2. Pick the set with the smallest number of stations that covers all 50 states

- This problem takes too much time to solve, since it's running time is _O(2^n)_, because there are _2^n_ stations
- You can make with 5 or 10 stations, but more than that and _there's no algorithm to solve it_

### Approximation Algorithms
- A greedy algorithm (not perfect, but good enough) way to solve it:
    - 1. Pick the station that covers most of the states that haven't been covered yet
    - 2. Repeat until are states are covered
    - In this case, the greedy algorithm runs in _O(n^2)_ time, where _n_ is the number of radio stations

- Approximation algorithms are judged by:
    - How fast they are
    - How close they are to the optimal solution

## Exercises
8.3:
> No

8.4:
> Yes

8.5:
> Yes

## NP-Complete Problems
- Problem that, for every new "vertice" the number of possible routes increments in factorial style
- Travel salesman problem. For every city you add, the number of routes grows exponentially
- 10 Cities == 10! == 3.628.800 possible routes
- Using a Greedy Algorithm, you pick a random city to start and the best path to next city and goes on and on. Maybe not the absolutely best route, but a good enough nonetheless

- How to tell if the problem in hand is a NP-Complete problem?
    - Algorithm runs fast with a couple of items but very slow with a large set of items
    - "All combinations of X", usually a NP-Complete problem
    - "Every possible solution", "Every possible route"
    - Involves a sequence or set and it's hard to solve
    - If you can rewrite your problem in the style of the travelling salesman or set-covering problem is definitely a NP-Complete

## Exercises
8.6:
> Yes

8.7:
> Yes

8.8:
> Yes