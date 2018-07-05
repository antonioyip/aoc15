package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	inputs, err := ioutil.ReadFile("day15.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	for _, line := range lines {
		var ing ingredient
		fmt.Sscanf(line, "%s: capacity %d, durability %d, flavor %d, texture %d, calories %d", &ing.name, &ing.capacity, &ing.durability, &ing.flavor, &ing.texture, &ing.calories)
	}
}
