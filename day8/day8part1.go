package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func codeSize(input string) int {
	return len(input)
}

func memorySize(input string) int {
	//return len(input) - strings.Count(input, "\\\\") - strings.Count(input, "\\\"") - 3*strings.Count(input, "\\x") - 2
	size := 0
	slash := false
	for _, char := range input {
		if slash {
			// handle slashes
			if char == 'x' {
				size -= 2 // exclude next two characters (e.g. \x27)
			} else if char == '"' {
				// do not count the quote character
			} else if char == '\\' {
				// do not count the second slash character
			} else {
				log.Fatal("Unexpected ", char, " after \\")
			}
			slash = false
		} else {
			if char == '\\' {
				slash = true
			}
			size++
		}
	}
	fmt.Println(size-2, " ", input)
	return size - 2 // assume all strings start and end with quote
}

func assert(expect int, result int) {
	if expect == result {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func main() {
	inputs, err := ioutil.ReadFile("day8.input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputs), "\n")
	totalCode := 0
	totalMemory := 0
	for _, line := range lines {
		totalCode += codeSize(line)
		totalMemory += memorySize(line)
	}
	fmt.Println(totalCode - totalMemory)
}
