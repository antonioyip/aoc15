package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"aoc15/graph"
)

func main() {
	inputs, err := ioutil.ReadFile("day13.input")
	if err != nil {
		panic(err)
	}

	// generate happinessGraph
	happinessGraph := graph.Graph{}
	lines := strings.Split(string(inputs), "\n")
	for _, line := range lines {
		var personA string
		var personB string
		var gainLose string
		var happiness int
		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &personA, &gainLose, &happiness, &personB)

		// strip the .
		personB = personB[:len(personB)-1]
		if gainLose == "lose" {
			happiness *= -1
		}
		happinessGraph.Add(personA, personB, happiness)
	}

	for personA, row := range happinessGraph {
		for personB, happiness := range row {
			fmt.Printf("%s : %s (%d)\n", personA, personB, happiness)
		}
	}
}
