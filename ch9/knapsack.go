package main

import "fmt"

type Item struct {
	Name   string
	Weight int
	Value  int
}

// Knapsack problem:
// You are a thief with a knapsick that can carry 4 kilos of goods
// In the store you are about to rob, there are 4 items with different values and weights (see items below)
// Question: What items should you steal to get the maximum amount of value? The items should fit your knapsack, of course.

var (
	// Items
	stereo *Item = &Item{
		Name:   "stereo",
		Weight: 4,
		Value:  3000,
	}
	laptop *Item = &Item{
		Name:   "laptop",
		Weight: 3,
		Value:  2000,
	}
	guitar *Item = &Item{
		Name:   "guitar",
		Weight: 1,
		Value:  1500,
	}
	iphone *Item = &Item{
		Name:   "iphone",
		Weight: 1,
		Value:  2000,
	}
)

func main() {
	// Creates a matrix with 4 rows (items above)
	// and 4 columns (knapsack size)
	cells := make([][][]*Item, 4)
	for i, _ := range cells {
		cells[i] = make([][]*Item, 4)
	}

	// We want a grid like this
	// 			1	2	3	4 	<- all possible weights for the knapsack
	// guitar   x   x   x   x   <- each row has a specific item
	// stereo   x   x   x   x
	// laptop   x   x   x   x
	// iphone   x   x   x   x

	for row, _ := range cells {
		item := getItemForRow(row)

		for col, _ := range cells[row] {
			knapsackWeight := col + 1

			// In the first row, we add our current item only if its weight is greater than the current sub knapsack weight
			// but we don't do crazy algorithms logic
			if row == 0 {
				if knapsackWeight >= item.Weight {
					cells[row][col] = append(cells[row][col], item)
				}
			} else {
				// In the second row onwards, things get harry!
				if knapsackWeight >= item.Weight {

					// Gets the sum of the Items' value on the cell above
					previousMaxValue := sumItemsValue(cells[row-1][col])

					// Do we still have space in your sub knapsack if we steal the current Item?
					remainingSpace := knapsackWeight - item.Weight

					// If we don't...
					if remainingSpace <= 0 {
						// Let's check the Item's value
						// If it is lesser than the cell right above us, let's keep that value,
						// since so far is the maximum value we can fit in our sub knapsack
						if item.Value < previousMaxValue {
							cells[row][col] = append(cells[row][col], cells[row-1][col]...)
						} else {
							// If not, let's still our current value
							cells[row][col] = append(cells[row][col], item)
						}
					} else {
						// BUT, if even after stealing the current Item we have space left, we can steal more!

						// First, we need to figure out how much space we have left in our sub knapsack
						spaceLeftInKnapsack := col - item.Weight

						// Second, we need to figure out the maximum value possible for that space
						// Luckily we already calculated that in our grid in the last row
						valueForSpaceLeft := sumItemsValue(cells[row-1][spaceLeftInKnapsack])

						// If we steal the current Item + fit all space left, we get the value below
						totalValue := item.Value + valueForSpaceLeft

						// Is this value bigger than the cell above us?
						if previousMaxValue <= totalValue {
							// If so, we add our current Item + the Items from the space left
							cells[row][col] = append(cells[row][col], item)
							cells[row][col] = append(cells[row][col], cells[row-1][spaceLeftInKnapsack]...)
						} else {
							// Otherwise we add the items from the cell above us
							cells[row][col] = append(cells[row][col], cells[row-1][col]...)
						}
					}
				} else {
					// Oh, if the current Item is heavier than our sub knapsack can support
					// just add the ones from the cell above
					cells[row][col] = append(cells[row][col], cells[row-1][col]...)
				}
			}
		}
	}

	// I'm looping over the cells again, _twice_, to get the the value and print the grid
	// I could have kept track of this information in the algorithm above, but I decided
	// to make it more clear to read ✨

	// Getting the bigger value!
	maxRow, maxCol, maxValue := getMaxiumValues(cells)

	// Print grid
	for row, _ := range cells {
		for col, items := range cells[row] {
			var s string
			for _, item := range items {
				s += item.Name + " "
			}

			if maxRow == row && maxCol == col {
				fmt.Printf("Cell(%d, %d) -> %s => THIS IS THE ANSWER! The value is: %d\n", row, col, s, maxValue)
			} else {
				fmt.Printf("Cell(%d, %d) -> %s\n", row, col, s)
			}
		}
	}
}

func getMaxiumValues(cells [][][]*Item) (maxRow, maxCol, maxValue int) {
	for row, _ := range cells {
		for col, items := range cells[row] {
			maxRow, maxCol, maxValue = row, col, sumItemsValue(cells[row][col])
			for _, _ = range items {
				m := sumItemsValue(items)
				if m > maxValue {
					maxRow, maxCol, maxValue = row, col, m
				}
			}
		}
	}
	return maxRow, maxCol, maxValue
}

// getItemForRow returns the Item for a specific row
func getItemForRow(row int) (item *Item) {
	switch row {
	case 0:
		item = guitar
	case 1:
		item = stereo
	case 2:
		item = laptop
	case 3:
		item = iphone
	}

	return item
}

// sumItemsValue return the sum of a slice of Items
func sumItemsValue(items []*Item) (sum int) {
	for _, item := range items {
		sum += item.Value
	}

	return sum
}
