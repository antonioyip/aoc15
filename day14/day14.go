package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Deer struct {
	name       string
	speedKps   int
	movingSec  int
	restingSec int
}

func main() {
	inputs, err := ioutil.ReadFile("day14.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	deers := []Deer{}
	for _, line := range lines {
		deer := Deer{}
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &deer.name, &deer.speedKps, &deer.movingSec, &deer.restingSec)
		deers = append(deers, deer)
	}

	maxDistance := 0
	for _, deer := range deers {
		maxDistance = max(maxDistance, distanceTraveled(deer, 2503))
	}
	fmt.Println(maxDistance)
}

func distanceTraveled(deer Deer, seconds int) int {
	quotient := seconds / (deer.movingSec + deer.restingSec)
	remainder := seconds % (deer.movingSec + deer.restingSec)

	distance := quotient*deer.speedKps*deer.movingSec + min(remainder, deer.movingSec)*deer.speedKps
	return distance
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
