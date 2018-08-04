package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// if incomingMin is equal, add incomingCount to currentCount
// if incomingMin is larger, ignore incomingCount
// if incomingMin is smaller, replace currentCount with incomingCount
// return final count, true min
func getCountAndMin(currentCount int, currentMin int, incomingCount int, incomingMin int) (int, int) {
	if incomingMin == currentMin {
		return currentCount + incomingCount, incomingMin
	} else if incomingMin < currentMin {
		return incomingCount, incomingMin
	} else {
		return currentCount, currentMin
	}
}

// returns total valid combinations and minimum total buckets used
func getValidMinimumCombinations(buckets []int, bucketsUsed int, minBuckets int, expectedSum int, currentSum int) (int, int) {
	if currentSum == expectedSum {
		return 1, bucketsUsed
	} else if currentSum > expectedSum {
		// adding more buckets won't help
		return 0, bucketsUsed
	} else if len(buckets) == 0 {
		// no more buckets
		return 0, bucketsUsed
	}

	sum := 0
	combinations := 0
	totalBuckets := 0

	// use current bucket
	combinations, totalBuckets = getValidMinimumCombinations(buckets[1:], bucketsUsed+1, minBuckets, expectedSum, currentSum+buckets[0])
	if combinations > 0 {
		if totalBuckets == minBuckets {
			sum += combinations
		} else if totalBuckets < minBuckets {
			// new minimum total buckets, reset count
			sum = combinations
			minBuckets = totalBuckets
		}
	}

	// skip current bucket
	combinations, totalBuckets = getValidMinimumCombinations(buckets[1:], bucketsUsed, minBuckets, expectedSum, currentSum)
	if combinations > 0 {
		if totalBuckets == minBuckets {
			sum += combinations
		} else if totalBuckets < minBuckets {
			// new minimum total buckets, reset count
			sum = combinations
			minBuckets = totalBuckets
		}
	}

	return sum, minBuckets
}

func sample() {
	buckets := []int{20, 15, 10, 5, 5}
	fmt.Println(getValidMinimumCombinations(buckets, 0, len(buckets), 25, 0))
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
	fmt.Println(getValidMinimumCombinations(buckets, 0, len(buckets), 150, 0))
}
