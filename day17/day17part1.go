package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// returns total valid combinations
func getValidCombinations(buckets []int, bucketIdx int, expectedSum int, currentSum int) int {
	if currentSum == expectedSum {
		return 1
	} else if currentSum > expectedSum {
		// adding more buckets won't help
		return 0
	} else if bucketIdx == len(buckets) {
		// no more buckets
		return 0
	}
	sum := 0
	// with current bucket
	sum += getValidCombinations(buckets, bucketIdx+1, expectedSum, currentSum+buckets[bucketIdx])
	// without current bucket
	sum += getValidCombinations(buckets, bucketIdx+1, expectedSum, currentSum)
	return sum
}

func sample() {
	buckets := []int{20, 15, 10, 5, 5}
	fmt.Println(getValidCombinations(buckets, 0, 25, 0))
}

func main() {
	inputs, err := ioutil.ReadFile("day17.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	buckets := make([]int, 0, len(lines))
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		buckets = append(buckets, val)
	}

	sample()
	fmt.Println(getValidCombinations(buckets, 0, 150, 0))
}
