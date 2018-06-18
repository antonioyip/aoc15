package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func codeSize(input string) int {
	return len(input)
}

func memorySize(input string) int {
	//return len(input) - strings.Count(input, "\\\\") - strings.Count(input, "\\\"") - 3*strings.Count(input, "\\x") - 2
	size := 0
	for _, char := range input {
		if char == '\\' {
			size++ // convert \ to \\
		} else if char == '"' {
			size++ // convert " to \"
		}
		size++
	}
	fmt.Println(size+2, " ", input)
	return size + 2 // add quotes to begining and end
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
	fmt.Println(totalMemory - totalCode)
}
