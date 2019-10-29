package main

import (
	"fmt"
	"math"
	"sort"
)

type Fruit struct {
	// banana or grapefruit
	Type string

	// 1 == orange
	// 10 == red
	Color float64

	// 1 == small
	// 10 == big
	Size float64
}

type Distance struct {
	Fruit    Fruit
	Distance float64
}

type Distances []Distance

func (d Distances) Len() int           { return len(d) }
func (d Distances) Less(i, j int) bool { return d[i].Distance < d[j].Distance }
func (d Distances) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func main() {
	fruits := []Fruit{
		{
			Type:  "orange",
			Color: 1,
			Size:  1,
		},
		{
			Type:  "orange",
			Color: 3,
			Size:  2,
		},
		{
			Type:  "orange",
			Color: 4,
			Size:  1,
		},
		{
			Type:  "orange",
			Color: 5,
			Size:  5,
		},
		{
			Type:  "orange",
			Color: 3,
			Size:  3,
		},
		{
			Type:  "orange",
			Color: 1,
			Size:  4,
		},
		{
			Type:  "orange",
			Color: 1,
			Size:  2,
		},
		{
			Type:  "orange",
			Color: 4,
			Size:  3,
		},
		{
			Type:  "grapefruit",
			Color: 7,
			Size:  5,
		},
		{
			Type:  "grapefruit",
			Color: 8,
			Size:  9,
		},
		{
			Type:  "grapefruit",
			Color: 10,
			Size:  5,
		},
		{
			Type:  "grapefruit",
			Color: 7,
			Size:  8,
		},
		{
			Type:  "grapefruit",
			Color: 8,
			Size:  5,
		},
		{
			Type:  "grapefruit",
			Color: 8,
			Size:  10,
		},
		{
			Type:  "grapefruit",
			Color: 9,
			Size:  7,
		},
		{
			Type:  "grapefruit",
			Color: 7,
			Size:  10,
		},
	}

	misteriousFruit := Fruit{
		Type:  "",
		Color: 8,
		Size:  10,
	}

	var distances Distances

	for _, item := range fruits {
		distance := Distance{
			Fruit:    item,
			Distance: math.Sqrt(math.Pow((item.Color-misteriousFruit.Color), 2) + math.Pow((item.Size-misteriousFruit.Size), 2)),
		}
		distances = append(distances, distance)
	}

	sort.Sort(distances)

	// the k for KNN
	k := 3
	for _, item := range distances[:k] {
		fmt.Println(item.Fruit)
	}
}
