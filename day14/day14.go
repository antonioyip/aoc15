package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type reindeer struct {
	name       string
	speedKps   int
	movingSec  int
	restingSec int
}

type racer struct {
	points   int
	distance int
}

func main() {
	inputs, err := ioutil.ReadFile("day14.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	deers := []reindeer{}
	for _, line := range lines {
		deer := reindeer{}
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &deer.name, &deer.speedKps, &deer.movingSec, &deer.restingSec)
		deers = append(deers, deer)
	}

	// part 1
	maxDistance := 0
	for _, deer := range deers {
		maxDistance = max(maxDistance, distanceTraveled(deer, 2503))
	}
	fmt.Println(maxDistance)

	// part 2
	maxDistance = 0
	racers := make([]racer, len(deers))
	for time := 1; time <= 2503; time++ {
		for deerIdx, deer := range deers {
			racers[deerIdx].distance = distanceTraveled(deer, time)
			maxDistance = max(maxDistance, racers[deerIdx].distance)
		}

		for deerIdx := range deers {
			if racers[deerIdx].distance == maxDistance {
				racers[deerIdx].points++
			}
		}
	}
	maxPoints := 0
	for _, racer := range racers {
		maxPoints = max(maxPoints, racer.points)
	}
	fmt.Println(maxPoints)
}

func distanceTraveled(deer reindeer, seconds int) int {
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
