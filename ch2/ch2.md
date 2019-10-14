# Selection Sort

## Arrays
- Uses contiguous memory, one address after another
- If want to add a new item but there is no slot, you need to move all the items to fit somewhere else
- One fix is to "hold space", allocate more space than needed to fit future allocation, but you might have unused space :( Or you even add more items than the pre allocated space, needing to move everything again
- Ordering is hard, need to rearrange everything. Same as deletion.

Pros:
- Good for reading random items, because you can find them with simple math

- *Reading Time*: _O(1)_
- *Insertion Time*: _O(n)_
- *Deletion Time*: _O(n)_

## Linked Lists
- Items can be anywhere in memory, each item has the address of the next item
- Doesn't need contiguous space, you can add items around and link them ⚡️
- Ordering is easier, just change pointers. Same as deletion.

Cons:
- Can't just read the last item, because you don't know where it is, you need to traverse the whole list
- Bad for jumping around

Pros:
- Good for reading all the items at once

- *Reading Time*: _O(n)_
- *Insertion Time*: _O(1)_
- *Deletion Time*: _O(1)_

## Exercise
2.1 
> A list!

2.2
> Linked list, there's no random access, things are added in the beginning of the queue and removed at the end. Also, Linked lists usually track their first and last element, so we just need to change the pointers around them!

2.3
> A sorted array! Because, since it's ordered, you can access items randomly doing simple math.

2.4
> If a new user is added you need to shift the whole array to add the user in the correct alphabetic place.

2.5
> Slower to search than arrays, since there's a latency of a linked list.
> Faster to insert than arrays, since the main array will never change, all the insertions will happen in the linked list.
> Faster to search than linked list, since you already have a separation of the alphabet.
> Same speed to insert as linked list, since the first operation is O(1), but the insertion happens in a list.

## Selection Sort
- Goes to every item in the list _n_ times and copy the smallest item to a new list, in order
- **Running time** is _O(n^2)_