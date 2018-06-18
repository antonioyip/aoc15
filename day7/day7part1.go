package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type gate struct {
	memoized bool
	value    uint16
	operator string
	valueA   string
	valueB   string
}

func (g *gate) getValue() uint16 {
	if g.memoized {
		return g.value
	}

	a64, errA := strconv.ParseUint(g.valueA, 10, 16)
	a16 := uint16(a64)
	if errA != nil {
		// not a signal (number), must be a wire
		a16 = circuit[g.valueA].getValue()
	}
	b64, errB := strconv.ParseUint(g.valueB, 10, 16)
	b16 := uint16(b64)
	if errB != nil {
		// not a signal (number), must be a wire
		b16 = circuit[g.valueB].getValue()
	}

	switch g.operator {
	case "AND":
		g.value = a16 & b16
	case "OR":
		g.value = a16 | b16
	case "LSHIFT":
		g.value = a16 << b16
	case "RSHIFT":
		g.value = a16 >> b16
	case "NOT":
		g.value = ^a16
	case "ASSIGN":
		g.value = a16
	default:
		panic(g.operator)
	}

	g.memoized = true
	return g.value
}

// circuit will be a map[string]gate
// wire name is the key
var circuit = map[string]*gate{}

func binaryOperator(token string) bool {
	var valueA, operator, valueB, destination string
	n, _ := fmt.Sscanf(token, "%s %s %s -> %s", &valueA, &operator, &valueB, &destination)
	if n != 4 {
		// ok; not all instructions are binary operators
		return false
	}
	circuit[destination] = &gate{
		false,
		0,
		operator,
		valueA,
		valueB,
	}
	return true
}

func unaryOperator(token string) bool {
	var operator, value, destination string
	n, _ := fmt.Sscanf(token, "%s %s -> %s", &operator, &value, &destination)
	if n != 3 {
		// ok; not all instructions are unary operators
		return false
	}
	circuit[destination] = &gate{
		false,
		0,
		operator,
		value,
		"0",
	}
	return true
}

func assignmentOperator(token string) bool {
	var value, destination string
	n, _ := fmt.Sscanf(token, "%s -> %s", &value, &destination)
	if n != 2 {
		// ok; not all instructions are assignment
		return false
	}
	circuit[destination] = &gate{
		false,
		0,
		"ASSIGN",
		value,
		"0",
	}
	return true
}

func main() {

	input, err := ioutil.ReadFile("day7.input")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(string(input), "\n")
	for _, instruction := range instructions {
		fmt.Println(instruction)
		if !(binaryOperator(instruction) || unaryOperator(instruction) || assignmentOperator(instruction)) {
			log.Fatal("Unexpected instruction: ", instruction)
		}
	}

	fmt.Println(circuit["a"].getValue())
}
