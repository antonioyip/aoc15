package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"aoc15/graph"
	"aoc15/permutations"
	"aoc15/stringList"
)

func main() {
	inputs, err := ioutil.ReadFile("day9.input")
	if err != nil {
		panic(err)
	}

	allCities := stringList.StringList{}
	distanceGraph := graph.Graph{}

	// parse the file
	// create the graph
	// create list of cities
	lines := strings.Split(string(inputs), "\n")
	for _, line := range lines {
		var source string
		var destination string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &source, &destination, &distance)

		allCities.AddUnique(source)
		allCities.AddUnique(destination)

		distanceGraph.Add(source, destination, distance)
		distanceGraph.Add(destination, source, distance)
	}

	minDistance := 999999
	maxDistance := 0
	allPermutations := permutations.GeneratePermutations(allCities)
	for _, aPermutation := range allPermutations {
		currentDistance := calcDistance(aPermutation, distanceGraph)
		fmt.Println(aPermutation, currentDistance)
		if minDistance > currentDistance {
			minDistance = currentDistance
		}
		if maxDistance < currentDistance {
			maxDistance = currentDistance
		}
	}
	fmt.Println("Minimum distance: ", minDistance)
	fmt.Println("Maximum distance: ", maxDistance)
}

func calcDistance(path []string, distances graph.Graph) int {
	totalDistance := 0
	for i := 1; i < len(path); i++ {
		totalDistance += distances[path[i-1]][path[i]]
	}
	return totalDistance
}
