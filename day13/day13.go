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
	inputs, err := ioutil.ReadFile("day13.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	happinessGraph := graph.Graph{}
	people := stringList.StringList{}

	// generate happinessGraph
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

		people.AddUnique(personA)
		people.AddUnique(personB)
		happinessGraph.Add(personA, personB, happiness)
	}

	maxHappiness := 0
	seatArrangements := permutations.GeneratePermutations(people)
	for _, arrangement := range seatArrangements {
		maxHappiness = max(maxHappiness, calcHappiness(arrangement, happinessGraph))
	}

	fmt.Println(maxHappiness)
}

func calcHappiness(people []string, happiness graph.Graph) int {
	sum := 0
	for i := 1; i < len(people); i++ {
		sum += happiness[people[i-1]][people[i]]
		sum += happiness[people[i]][people[i-1]]
	}
	sum += happiness[people[0]][people[len(people)-1]]
	sum += happiness[people[len(people)-1]][people[0]]
	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
