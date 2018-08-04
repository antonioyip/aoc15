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
	auntItems := 3
	for _, line := range lines {
		id := 0
		items := make([]string, auntItems)
		qty := make([]int, auntItems)
		fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d", &id, &items[0], &qty[0], &items[1], &qty[1], &items[2], &qty[2])

		aunt := person{id, map[string]int{}}
		for i, item := range items {
			// remove the trailing : from items when populating item map
			aunt.items[item[:len(item)-1]] = qty[i]
		}
		aunts = append(aunts, aunt)
	}

	mfcsamTape := map[string]int{}
	mfcsamTape["children"] = 3
	mfcsamTape["cats"] = 7
	mfcsamTape["samoyeds"] = 2
	mfcsamTape["pomeranians"] = 3
	mfcsamTape["akitas"] = 0
	mfcsamTape["vizslas"] = 0
	mfcsamTape["goldfish"] = 5
	mfcsamTape["trees"] = 3
	mfcsamTape["cars"] = 2
	mfcsamTape["perfumes"] = 1

	// part 1
	for _, aunt := range aunts {
		correctAunt := true
		for itemName, itemCount := range aunt.items {
			if mfcsamTape[itemName] != itemCount {
				correctAunt = false
			}
		}
		if correctAunt {
			fmt.Println(aunt)
		}
	}

	// part 2
	for _, aunt := range aunts {
		correctAunt := true
		for itemName, itemCount := range aunt.items {
			switch itemName {
			case "cats", "trees":
				if mfcsamTape[itemName] >= itemCount {
					correctAunt = false
				}
			case "pomeranians", "goldfish":
				if mfcsamTape[itemName] <= itemCount {
					correctAunt = false
				}
			default:
				if mfcsamTape[itemName] != itemCount {
					correctAunt = false
				}
			}
		}
		if correctAunt {
			fmt.Println(aunt)
		}
	}
}
