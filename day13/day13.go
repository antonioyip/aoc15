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
	inputs, err := ioutil.ReadFile("day13.input")
	if err != nil {
		panic(err)
	}

	// generate happinessGraph
	happinessGraph := graph{}
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
		happinessGraph.add(personA, personB, happiness)
	}

	for personA, row := range happinessGraph {
		for personB, happiness := range row {
			fmt.Printf("%s : %s (%d)\n", personA, personB, happiness)
		}
	}
}
