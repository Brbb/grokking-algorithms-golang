# Breadth-first Search
- Algorithm to find the shortest distance between two things, i.e:
    - Checkers AI to calculate fewest moves to victory
    - Spell checker
    - Find someone closer in my network
- Answers two questions:
    - Is there a path from node A to node B?
    - What is the shortest path from node A to node B?
        - In BFS, you do:
            - Get all your neighbors and put in a Queue
            - Initialize a _searched_ list

            - Go through the Queue, _popping_ one by one (1st degree connections)
                - If they were already searched, break the loop (you don't wanna search the same neighbor twice or end up in an infinite loop).

                - If they answer your question, good!
                - If not, _push_ all their friends (2st degree connections)
                - Add the neighbor in the _searched_ list
            
            - This way you always search your 1st degree first, 2nd degree second etc
            - You shouldn't search any 2nd connections before exhausting all 1st connections
- Running time: _O(number of people + number of edges)_ or _O(V+E)_.
    - Number of people because because of the _searched_ array
    - Number of edges because if you go through all the edges, well, it will be all the edges
    

## Introduction to Graphs
- Graphs models a set of connections
- Graphs are composed of Nodes and Edges (circles and arrows ;)
- A Node can directly connected to any Nodes, they are called its neighbors
- Like:
    JOJO     ->    THEODORE    ->      MAX
    (node) (edge)  (node)    (edge)   (node)

    - Here, Jojo and Theodore are neighbors
    - But Jojo and Max aren't

    - This is also a _directed graph_, where arrows points to only one direction
    - In an _undirected graph_, there are no arrows, just lines, both nodes connect to themselves

- A Graph can be _sorted_, in a way, if Edges _depends_ on other Edges. This is called a _topological sort_, a way to make an ordered list from a Graph.

- A Graph cam also be a _tree_, when no edges ever point back. Think about a family tree, where the edges just do downwards (you can't be your dad's dad).

## Queue
- Data structure implementing a _FIFO_, First In First Out
- Has usually two operations:
    - _Enqueue_ or _Push_: add something to the end of the Queue
    - _Dequeue_ or _Pop_: pops something from the front of the Queue

## Exercises
6.1
> 2

6.2
> 2

6.3
> A = invalid  
> B = valid  
> C = invalid

6.4
> ...

6.5
> A, C