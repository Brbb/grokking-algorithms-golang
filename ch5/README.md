# Hash Tables
- A data structure to hold value/key pairs
- At best case, it will have _O(1)_ (constant time) time for inserting, searching and deleting
- At worst, it will be _O(n)_ (linear time)

- Pass a key through the a _hash function_, it will return return a number, the same number for the same input (in most cases)
- Use this number to find a place in an array to store the key's value
- When you want to find the value of the key, go through the _hash function_ and you will have the exact place! _O(n1)_

## Hash Functions
- The hash function is one of the most important parts of the Hash Table
- It needs:
    - To be consistent: always return the same value for the same input
    - Should map input to different inputs, is no good if always returns "1"
- The hash function knows the array's size and only return valid indexes numbers, so u don't put values in an inexistent index

## Collisions
- A hash function would be perfect if it returns a consistently different output for every different input, but that's hard, almost impossible
- If it returns the same output for two different inputs, we have a _collision_
- If there's a _collision_, the spot in the array becomes a linked list, and the values are added there. When this happens, the search time on that index becomes _O(n)_, the linked list running time
- A good hash function avoid collisions, to avoid long linked lists, so all items are balanced in the underlying store array

- **Load factor**: the number of occupied numbers in the hash table array, calculated by _number of items in the hash table_ / _total number of slots_. Generally, languages will resize the array when the load factor is a bit before one, which means copying everything to a new array, but in average hash tables still takes _O(1)_ even with resizing

## Exercises 
5.5
> D

5.6
> B

5.7
> D

## Resources
http://shlegeris.com/2017/01/06/hash-maps
