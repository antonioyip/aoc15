package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type graph map[string]map[string]int

type stringList []string

func (g graph) add(key1 string, key2 string, value int) {
	if _, ok := g[key1]; !ok {
		g[key1] = make(map[string]int)
	}
	if _, ok := g[key2]; !ok {
		g[key2] = make(map[string]int)
	}
	g[key1][key2] = value
	g[key2][key1] = value
}

func (array stringList) contains(value string) bool {
	for _, entry := range array {
		if entry == value {
			return true
		}
	}
	return false
}

func main() {
	inputs, err := ioutil.ReadFile("day9.input")
	if err != nil {
		panic(err)
	}

	allCities := stringList{}
	distanceGraph := graph{}

	// parse the file
	// create the graph
	// create list of cities
	lines := strings.Split(string(inputs), "\n")
	for _, line := range lines {
		var source string
		var destination string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &source, &destination, &distance)

		if !allCities.contains(source) {
			allCities = append(allCities, source)
		}
		if !allCities.contains(destination) {
			allCities = append(allCities, destination)
		}

		distanceGraph.add(source, destination, distance)
	}

	minDistance := 999999
	maxDistance := 0
	allPermutations := permutation(len(allCities), allCities)
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

// Heap's algorithm
func permutation(length int, values []string) [][]string {
	result := [][]string{}
	if length == 1 {
		// force a deep copy
		temp := make([]string, len(values))
		copy(temp, values)
		result = append(result, temp)
	} else {
		for i := 0; i < length-1; i++ {
			result = append(result, permutation(length-1, values)...)
			if length%2 == 0 {
				values[i], values[length-1] = values[length-1], values[i]
			} else {
				values[0], values[length-1] = values[length-1], values[0]
			}
		}
		result = append(result, permutation(length-1, values)...)
	}
	return result
}

func calcDistance(path []string, distances graph) int {
	totalDistance := 0
	for i := 1; i < len(path); i++ {
		totalDistance += distances[path[i-1]][path[i]]
	}
	return totalDistance
}
