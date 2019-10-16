# Dijkstra Algorithm
- Finds the shortest path from a start node to all the Nodes in a Graph
- Works in a weighted graph, where the Edges has a weight (or cost)
- Doesn't work if an Edge has a negative weight
- Running time is _O(E log V)_, where _E_ is the total number of Edges and _V_ is the total number of Vertices (Nodes)

## How does it work?
- The code on `dijkstra.go` is very well documented, please take a look there, but here's a simple version

- Have a Graph with weighted Edges (with Parent, Child and Cost attributes)
- Pick a Node to be your `startNode`, let's say it is `A`
- Create a `costTable`, a table specifying the cost from the `startNode` to all the other Nodes
    - It will look like this in the first run:
    ```
    NODE  COST
    A     0
    B     Inf
    C     Inf
    D     Inf
    ```
    - The `sourceNode` always get the lowest cost, in this case being `0`
    - All the distance to other Nodes are yet unknown, we mark it as `Infinity`
- Create an empty list of Nodes called `visited` to mark the Nodes as visited, because we only process each Node once

- While we didn't visit all the Nodes:
    - Get the `closestNonVisitedNode`: a Node from the `costTable` with the lower cost that we didn't visit
    - Mark is as `visited`
    - For each `Edge` of this Node (it's neighbors):
        - Get the cost to get from the picked Node to the Edge Child: this is the sum of the `costTable[pickedNode]` + `Edge.Cost`
        - If this cost is _lesser_ than the cost already in the table for that Edge Child (neighbor – `tableCost[edge.Child]`), **update the cost of the Edge.Child in the costTable with the calculated cost**

- After visiting all the Nodes, `costTable` will have the shortest path to all of them

- And _scene_.

## Exercises
> See dijkstra_exercises.go

## Resources
https://www.youtube.com/watch?v=_lHSawdgXpI – very quick demonstration, nice to start
https://brilliant.org/wiki/dijkstras-short-path-finder/ - loved it, the `dijkstra.go` is based on it
http://web.cecs.pdx.edu/~sheard/course/Cs163/Doc/Graphs.html
https://rosettacode.org/wiki/Dijkstra's_algorithm
