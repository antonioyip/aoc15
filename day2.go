package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// read input line by line
	file, err := os.Open("day2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, isPrefix, err := reader.ReadLine()
	totalWrapping := 0
	totalRibbon := 0
	for err == nil && !isPrefix {
		dimensions := getDimensions(string(line))
		area := calculateWrappingArea(dimensions)
		length := calculateRibbonLength(dimensions)
		totalWrapping += area
		totalRibbon += length
		fmt.Printf("%s\n%d\n", line, length)
		line, isPrefix, err = reader.ReadLine()
	}

	if isPrefix {
		log.Fatal("Buffer is too small")
		return
	}

	if err != io.EOF {
		log.Fatal(err)
	}

	fmt.Printf("Wrapping area: %d\n", totalWrapping)
	fmt.Printf("Ribbon length: %d\n", totalRibbon)
}

/// input - dimensions in the format LxWxH
/// return an array of integers containing L, W, H
func getDimensions(lxwxh string) []int {
	tokens := strings.Split(lxwxh, "x")
	if len(tokens) != 3 {
		log.Fatal("Invalid dimensions for ", lxwxh)
	}

	dimensions := make([]int, 3)
	for idx, token := range tokens {
		i, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		dimensions[idx] = i
	}
	return dimensions
}

/// dimensions - array of integers containing L, W, H
/// return area needed for wrapping
func calculateWrappingArea(dimensions []int) int {
	val1 := dimensions[0] * dimensions[1]
	val2 := dimensions[0] * dimensions[2]
	val3 := dimensions[1] * dimensions[2]

	min := val1
	if val2 < min {
		min = val2
	}
	if val3 < min {
		min = val3
	}

	return 2*val1 + 2*val2 + 2*val3 + min
}

/// dimensions - array of integers containing L, W, H
/// return length of ribbons
func calculateRibbonLength(dimensions []int) int {
	length := dimensions[0] * dimensions[1] * dimensions[2]
	val1 := dimensions[0] + dimensions[1]
	val2 := dimensions[0] + dimensions[2]
	val3 := dimensions[1] + dimensions[2]

	min := val1
	if val2 < min {
		min = val2
	}
	if val3 < min {
		min = val3
	}

	return length + 2*min
}
