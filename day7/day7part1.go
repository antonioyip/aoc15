package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// circuit will be a map[string]int
// wire name is the key, value is the signal
var circuit = map[string]uint16{}

func binaryOperator(token string) bool {
	var valueA, operator, valueB, destination string
	n, _ := fmt.Sscanf(token, "%s %s %s -> %s", &valueA, &operator, &valueB, &destination)
	if n != 4 {
		// ok; not all instructions are binary operators
		return false
	}

	a64, errA := strconv.ParseUint(valueA, 10, 16)
	a16 := uint16(a64)
	if errA != nil {
		// not a signal (number), must be a wire
		a16 = circuit[valueA]
	}
	b64, errB := strconv.ParseUint(valueB, 10, 16)
	b16 := uint16(b64)
	if errB != nil {
		// not a signal (number), must be a wire
		b16 = circuit[valueB]
	}

	switch operator {
	case "AND":
		circuit[destination] = a16 & b16
	case "OR":
		circuit[destination] = a16 | b16
	case "LSHIFT":
		circuit[destination] = a16 << b16
	case "RSHIFT":
		circuit[destination] = a16 >> b16
	default:
		panic(operator)
	}
	fmt.Println(destination, " = ", circuit[destination])
	fmt.Println("binary")
	return true
}

func unaryOperator(token string) bool {
	var operator, value, destination string
	n, _ := fmt.Sscanf(token, "%s %s -> %s", &operator, &value, &destination)
	if n != 3 {
		// ok; not all instructions are unary operators
		return false
	}

	i64, errA := strconv.ParseUint(value, 10, 16)
	i16 := uint16(i64)
	if errA != nil {
		// not a signal (number), must be a wire
		i16 = circuit[value]
	}

	switch operator {
	case "NOT":
		circuit[destination] = ^i16
	default:
		panic(operator)
	}
	fmt.Println(destination, " = ", circuit[destination])
	fmt.Println("unary")
	return true
}

func assignmentOperator(token string) bool {
	var value, destination string
	n, _ := fmt.Sscanf(token, "%s -> %s", &value, &destination)
	if n != 2 {
		// ok; not all instructions are assignment
		return false
	}

	i64, errA := strconv.ParseUint(value, 10, 16)
	i16 := uint16(i64)
	if errA != nil {
		// not a signal (number), must be a wire
		i16 = circuit[value]
	}

	circuit[destination] = i16
	fmt.Println(destination, " = ", circuit[destination])
	fmt.Println("assignment")
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

	fmt.Println(circuit["a"])
}
