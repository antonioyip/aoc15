package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type person struct {
	id    int
	items map[string]int
}

func main() {
	inputs, err := ioutil.ReadFile("day16.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	aunts := []person{}
	itemCount := 3
	for _, line := range lines {
		id := 0
		items := make([]string, itemCount)
		qty := make([]int, itemCount)
		fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d", &id, &items[0], &qty[0], &items[1], &qty[1], &items[2], &qty[2])

		aunt := person{id, map[string]int{}}
		for i, item := range items {
			// remove the trailing : from items when populating item map
			aunt.items[item[:len(item)-1]] = qty[i]
		}
		aunts = append(aunts, aunt)
	}

	for _, aunt := range aunts {
		fmt.Println(aunt)
	}
}
