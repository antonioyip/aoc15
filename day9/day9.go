package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type edge struct {
	a string
	b string
}

func (e edge) toString() string {
	return e.a + " -> " + e.b
}

var graph = map[edge]int{}

var cities = map[string]bool{}

func main() {
	inputs, err := ioutil.ReadFile("day9.input")
	if err != nil {
		panic(err)
	}

	// parse the file
	// create the graph object
	// create a list of unique cities
	lines := strings.Split(string(inputs), "\n")
	for _, line := range lines {
		var source string
		var destination string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &source, &destination, &distance)
		graph[edge{source, destination}] = distance
		graph[edge{destination, source}] = distance

		if _, ok := cities[source]; !ok {
			cities[source] = true
		}

		if _, ok := cities[destination]; !ok {
			cities[destination] = true
		}
	}

	// print out the graph entries
	// for e, v := range graph {
	// 	fmt.Printf("%d: %s\n", v, e.toString())
	// }

	for c1 := range cities {
		for c2 := range cities {
			if c1 == c2 {
				continue
			}
			fmt.Printf("%s -> %s\n", c1, c2)
		}
	}

}
