package main

import (
	"reflect"
	"testing"
	"time"
)

type Class struct {
	Name  string
	Start time.Time
	End   time.Time
}

// getBestClasses solves the classroom scheduling problem. Given a set of overlapping *but ordered*
// classes, find the set of classes that uses the best of time
func getBestClasses(classes []Class) []Class {
	bestClasses := []Class{}
	// get class that ends the soonest, since it's ordered, it's the first element
	bestClasses = append(bestClasses, classes[0])

	for _, c := range classes[1:] {
		// pick a class that starts after the first class ends
		if c.Start.Equal(bestClasses[len(bestClasses)-1].End) || c.Start.After(bestClasses[len(bestClasses)-1].End) {
			bestClasses = append(bestClasses, c)
		}
	}

	return bestClasses
}

func Test_ClassRoom(t *testing.T) {
	nineAM, _ := time.Parse(time.Kitchen, "9:00AM")
	tenAM, _ := time.Parse(time.Kitchen, "10:00AM")
	ninehalfAM, _ := time.Parse(time.Kitchen, "9:30AM")
	tenhalfAM, _ := time.Parse(time.Kitchen, "10:30AM")
	elevenAM, _ := time.Parse(time.Kitchen, "11:00AM")
	elevenhalfAM, _ := time.Parse(time.Kitchen, "11:30AM")
	twelvePM, _ := time.Parse(time.Kitchen, "12:00PM")

	classes := []Class{
		{
			Name:  "art",
			Start: nineAM,
			End:   tenAM,
		},
		{
			Name:  "eng",
			Start: ninehalfAM,
			End:   tenhalfAM,
		},
		{
			Name:  "math",
			Start: tenAM,
			End:   elevenAM,
		},
		{
			Name:  "cs",
			Start: tenhalfAM,
			End:   elevenhalfAM,
		},
		{
			Name:  "music",
			Start: elevenAM,
			End:   twelvePM,
		},
	}

	classesToHold := getBestClasses(classes)
	bestClassesToHold := []Class{
		{
			Name:  "art",
			Start: nineAM,
			End:   tenAM,
		},
		{
			Name:  "math",
			Start: tenAM,
			End:   elevenAM,
		},
		{
			Name:  "music",
			Start: elevenAM,
			End:   twelvePM,
		},
	}

	if !reflect.DeepEqual(classesToHold, bestClassesToHold) {
		t.Errorf("expected %v, got %v", bestClassesToHold, classesToHold)
	}
}
