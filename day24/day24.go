package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs, err := ioutil.ReadFile("day24.input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputs), "\n")
	packages := make([]int, len(lines))
	totalWeight := 0
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		packages[i] = num
		totalWeight += num
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packages)))

	state := State{[][]int{}, totalWeight / 3, len(packages)}
	FindGroups(packages, []int{}, &state)
	fmt.Println(GetMinQuantumEntanglement(packages, state))

	state.expectedWeight = totalWeight / 4
	FindGroups(packages, []int{}, &state)
	fmt.Println(GetMinQuantumEntanglement(packages, state))
}

func GetMinQuantumEntanglement(packages []int, state State) int {
	minProduct := math.MaxInt64
	for _, arrangement := range state.validArrangements {
		if len(arrangement) != state.minArrangementSize {
			continue
		}
		if !ValidateArrangement(packages, arrangement) {
			continue
		}
		product := 1
		for _, num := range arrangement {
			product *= num
		}
		if minProduct > product {
			minProduct = product
		}
	}
	return (minProduct)
}

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

type State struct {
	validArrangements  [][]int
	expectedWeight     int
	minArrangementSize int
}

func FindGroups(packages []int, currentArrangement []int, state *State) {
	if len(currentArrangement) > state.minArrangementSize {
		return
	}
	for i, weight := range packages {
		newArrangement := Join(currentArrangement, weight)
		newArrangementWeight := Sum(newArrangement)
		if newArrangementWeight == state.expectedWeight {
			state.validArrangements = append(state.validArrangements, newArrangement)
			if state.minArrangementSize > len(newArrangement) {
				state.minArrangementSize = len(newArrangement)
			}
		} else if newArrangementWeight < state.expectedWeight {
			FindGroups(packages[i+1:], newArrangement, state)
		}
	}
}

func FindOneGroup(packages []int, currentArrangement []int, state *State) bool {
	for i, weight := range packages {
		newArrangement := append(currentArrangement, weight)
		newArrangementWeight := Sum(newArrangement)
		if newArrangementWeight == state.expectedWeight {
			state.validArrangements = append(state.validArrangements, newArrangement)
			return true
		} else if newArrangementWeight < state.expectedWeight {
			if FindOneGroup(packages[i+1:], newArrangement, state) {
				return true
			}
		}
	}
	return false
}

func Join(array []int, newNumber int) []int {
	newArray := make([]int, len(array)+1)
	for i, num := range array {
		newArray[i] = num
	}
	newArray[len(array)] = newNumber
	return newArray
}

func Contains(array []int, number int) bool {
	for _, num := range array {
		if num == number {
			return true
		}
	}
	return false
}

func Remove(source []int, exclude []int) []int {
	destination := make([]int, 0, len(source))
	for _, num := range source {
		if !Contains(exclude, num) {
			destination = append(destination, num)
		}
	}
	return destination
}

// TODO - this validate 3 groups; validate arrangement for 4 groups
func ValidateArrangement(packages []int, arrangement []int) bool {
	state := State{[][]int{}, Sum(arrangement), 0}
	packages = Remove(packages, arrangement)
	if FindOneGroup(packages, []int{}, &state) {
		// fmt.Println(Sum(arrangement), Sum(state.validArrangements[0]), Sum(Remove(packages, state.validArrangements[0])), arrangement, state.validArrangements[0], Remove(packages, state.validArrangements[0]))
		return true
	}
	return false
}
