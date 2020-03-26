# Dynamic Programming
_Very hard concept to grasp. Use this notes with the book_ 🤓

- A technique to solve a hard problem by breaking it upt into sub problems and solving those sub problems first

- Every dynamic programming algorithm starts with a grid

- You can keep adding items to the grid, the value will never go down (each iteration calculates the maximum possible value for that line/column)

- The order of the rows (sub problems) don't matter

- Depending on the problem, you'll need to fill the grid column-wise instead of row-wise, but sometimes it doesn't matter

- If you want to add a _granular_ sub problem, you need to recalculate. Let's say you are accounting for points and they are just integers (1, 2, 3, 5, 10), but if you want to add a 0.5 point you need to recalculate, since you need to have the solutions for the 0.5, 1, 1.5 etc problems

- With dynamic programming, you either get an item during the problem solving or not, so you can't take half _item_. But you can mix it up with a greed algorithm: for each item, take as much as you can from the most valuable one and keep going

- Dynamic programming only works when each sub problem is discrete – when it doesn't depend on other sub problems

- **Dynamic programming is useful when you are trying to optimize something given a constraint**

- You can use dynamic programming when the problem can be broken into discrete sub problems when they don't depend on each other

- **Thinking about a dynamic programming problem**
    - Involves a grid
    - The value of the cells are what we are trying to optimize (value of the goods in the knapsack problem)
    - Each cell is a sub problem

- There's no single formula to calculate a dynamic programming algorithm, it depends on the problem