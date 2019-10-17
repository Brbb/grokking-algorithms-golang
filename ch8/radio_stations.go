package main

import "fmt"

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	finalStations := []string{}

	for len(statesNeeded) > 0 {
		var bestStation string
		var statesCovered []string

		for station, states := range stations {
			covered := []string{}

			for _, sn := range statesNeeded {
				for _, st := range states {
					if sn == st {
						covered = append(covered, sn)
					}
				}
			}

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		for _, sc := range statesCovered {
			for i, sn := range statesNeeded {
				if sc == sn {
					statesNeeded = append(statesNeeded[:i], statesNeeded[i+1:]...)
				}
			}
		}

		finalStations = append(finalStations, bestStation)

	}

	fmt.Println(finalStations)
}
