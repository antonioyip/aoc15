package main

import (
	"fmt"
	"strings"
)

func main() {
	test()
	input := "cqjxjnds"
	for i := 0; i < 2; i++ {
		input = nextPassword(input)
		fmt.Println(input)
	}

}

func increasingLetters(input string) bool {
	for i := 2; i < len(input); i++ {
		if input[i] == input[i-1]+1 && input[i] == input[i-2]+2 {
			return true
		}
	}
	return false
}

func validLetters(input string, invalidChars string) bool {
	return !strings.ContainsAny(input, invalidChars)
}

func twoPair(input string) bool {
	pairs := 0
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			pairs++
			i++ // do not overlap
		}
	}

	return pairs > 1
}

func nextCandidate(password string) string {
	for i := len(password) - 1; i >= 0; i-- {
		char := password[i] + 1
		if char > 'z' {
			char = 'a'
			password = password[:i] + string(char) + password[i+1:]
		} else {
			password = password[:i] + string(char) + password[i+1:]
			break
		}
	}

	return password
}

func assert(expected bool, result bool) {
	if expected == result {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func test() {
	assert(true, nextCandidate("xx") == "xy")
	assert(true, nextCandidate("xy") == "xz")
	assert(true, nextCandidate("xz") == "ya")
	assert(true, nextCandidate("ya") == "yb")
	assert(true, nextCandidate("azz") == "baa")
	assert(true, increasingLetters("hijklmmn"))
	assert(false, validLetters("hijklmmn", "i | o | l"))
	assert(false, increasingLetters("abbceffg"))
	assert(true, twoPair("abbceffg"))
	assert(false, twoPair("abbcegjk"))
	assert(true, nextPassword("abcdefgh") == "abcdffaa")
	assert(true, nextPassword("ghijklmn") == "ghjaabcc")
}

func nextPassword(input string) string {
	for {
		input = nextCandidate(input)
		if increasingLetters(input) &&
			validLetters(input, "i | o | l") &&
			twoPair(input) {
			return input
		}
	}
}
