package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type graph map[string]map[string]int

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

func main() {
	inputs, err := ioutil.ReadFile("day9.input")
	if err != nil {
		panic(err)
	}

	citiesVisited := make(map[string]bool)
	distanceGraph := graph{}

	// parse the file
	// create the graph
	lines := strings.Split(string(inputs), "\n")
	for _, line := range lines {
		var source string
		var destination string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &source, &destination, &distance)

		citiesVisited[source] = false
		citiesVisited[destination] = false

		distanceGraph.add(source, destination, distance)
	}

	// print out city list
	for city, visited := range citiesVisited {
		fmt.Printf("%s %t\n", city, visited)
	}

	// print out the graph entries
	for city1, row := range distanceGraph {
		for city2, distance := range row {
			fmt.Printf("%d: %s -> %s\n", distance, city1, city2)
		}
	}

	totalDistance := 0
	for city := range citiesVisited {
		citiesVisited[city] = true
		thisDistance := tsp(distanceGraph, city, citiesVisited)
		if totalDistance < thisDistance {
			totalDistance = thisDistance
		}
		citiesVisited[city] = false
	}

	fmt.Printf("Total distance %d\n", totalDistance)
}

func tsp(distanceGraph graph, currentCity string, visitedCities map[string]bool) int {
	totalDistance := 0
	for city, visited := range visitedCities {
		if !visited {
			visitedCities[city] = true
			currentDistance := tsp(distanceGraph, city, visitedCities) + distanceGraph[currentCity][city]
			if totalDistance < currentDistance {
				totalDistance = currentDistance
			}
			visitedCities[city] = false
		}
	}
	return totalDistance
}
