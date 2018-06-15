package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func threeVowels(input string) bool {
	total := 0
	for _, char := range input {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			total++
			if total == 3 {
				return true
			}
		}
	}
	return false
}

func twoConsecutive(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}

func validPatterns(input string) bool {
	return !regexp.MustCompile("ab|cd|pq|xy").MatchString(input)
}

func niceString(input string) bool {
	return twoConsecutive(input) && threeVowels(input) && validPatterns(input)
}

func assert(expected bool, result bool) {
	if result == expected {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func main() {
	assert(true, twoConsecutive("aa"))
	assert(false, twoConsecutive("ab"))
	assert(true, threeVowels("aei"))
	assert(false, threeVowels("oub"))
	assert(false, validPatterns("xy"))
	assert(true, validPatterns("ca"))
	assert(true, niceString("ugknbfddgicrmopn"))
	assert(true, niceString("aaa"))
	assert(false, niceString("jchzalrnumimnmhp"))
	assert(false, niceString("jchzalrnumimnmhp"))
	assert(false, niceString("dvszwmarrgswjxmb"))

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
