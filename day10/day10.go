package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(input string) string {
	var outputBuilder strings.Builder
	currentRune := input[0]
	currentCounter := 1
	for i := 1; i < len(input); i++ {
		if currentRune != input[i] {
			outputBuilder.WriteString(strconv.Itoa(currentCounter))
			outputBuilder.WriteString(string(currentRune))
			// reset
			currentRune = input[i]
			currentCounter = 0
		}
		currentCounter++
	}
	outputBuilder.WriteString(strconv.Itoa(currentCounter))
	outputBuilder.WriteString(string(currentRune))
	return outputBuilder.String()
}

func main() {
	input := "3113322113"
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	fmt.Println(len(input))
	for i := 40; i < 50; i++ {
		input = lookAndSay(input)
	}
	fmt.Println(len(input))
}
