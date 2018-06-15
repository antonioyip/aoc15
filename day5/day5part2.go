package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func twoPair(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if strings.Contains(input[i+2:], input[i:i+2]) {
			return true
		}
	}
	return false
}

func repeatWithGap(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false
}

func niceString(input string) bool {
	return repeatWithGap(input) && twoPair(input)
}

func assert(expected bool, result bool) {
	if result == expected {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func main() {
	assert(true, repeatWithGap("xyx"))
	assert(true, repeatWithGap("aaa"))
	assert(false, repeatWithGap("kdf"))
	assert(false, twoPair("aa"))
	assert(false, twoPair("aaa"))
	assert(true, twoPair("aaaa"))
	assert(true, twoPair("aaxxlmnxx"))
	assert(true, niceString("qjhvhtzxzqqjkmpb"))
	assert(true, niceString("xxyxx"))
	assert(false, niceString("uurcxstgmygtbstg"))
	assert(false, niceString("ieodomkazucvgmuy"))

	input, err := ioutil.ReadFile("day5.input")
	if err != nil {
		panic(err)
	}

	total := 0
	inputs := strings.Split(string(input), "\n")
	for _, str := range inputs {
		if niceString(str) {
			total++
		}
	}
	fmt.Println("Nice strings: ", total)
}
