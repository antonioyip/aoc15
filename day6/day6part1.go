package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type grid [1000][1000]bool

var lights grid
var testLights grid

type box struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func newBox(values []string) *box {
	if len(values) != 4 {
		panic("Invalid coordinates count")
	}

	numbers := make([]int, 0, 4)
	for _, val := range values {
		number, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	return &box{numbers[0], numbers[1], numbers[2], numbers[3]}
}

// sort.Ints(b.x[:])
// sort.Ints(b.y[:])

func (g *grid) turnOn(b box) {
	for x := b.x1; x <= b.x2; x++ {
		for y := b.y1; y <= b.y2; y++ {
			g[x][y] = true
		}
	}
}

func (g *grid) turnOff(b box) {
	for x := b.x1; x <= b.x2; x++ {
		for y := b.y1; y <= b.y2; y++ {
			g[x][y] = false
		}
	}
}

func (g *grid) toggle(b box) {
	for x := b.x1; x <= b.x2; x++ {
		for y := b.y1; y <= b.y2; y++ {
			g[x][y] = !g[x][y]
		}
	}
}

func (g *grid) countOn() int {
	totalOn := 0
	for _, row := range g {
		for _, col := range row {
			if col {
				//fmt.Printf("%d ", 1)
				totalOn++
			} else {
				//fmt.Printf("%d ", 0)
			}
		}
		//fmt.Printf("\n")
	}
	return totalOn
}

func assert(expected int, result int) {
	if result == expected {
		fmt.Println("Success")
	} else {
		fmt.Printf("Failure %d vs %d\n", result, expected)
	}
}

func main() {

	testLights.turnOn(box{4, 4, 5, 5})
	assert(4, testLights.countOn())
	testLights.toggle(box{4, 4, 4, 5})
	assert(2, testLights.countOn())

	input, err := ioutil.ReadFile("day6.input")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(input), "\n")
	re := regexp.MustCompile("(\\d+),(\\d+) through (\\d+),(\\d+)")
	for _, instruction := range inputs {
		matches := re.FindStringSubmatch(instruction)
		b := newBox(matches[1:])

		if strings.Contains(instruction, "turn on") {
			lights.turnOn(*b)
		} else if strings.Contains(instruction, "turn off") {
			lights.turnOff(*b)
		} else if strings.Contains(instruction, "toggle") {
			lights.toggle(*b)
		} else {
			log.Fatal("Unexpected instruction ", instruction)
		}
	}
	fmt.Println(lights.countOn())
}
