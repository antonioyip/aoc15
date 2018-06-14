package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
		totalWrapping += calculateWrappingArea(dimensions)
		totalRibbon += calculateRibbonLength(dimensions)
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
	dimensions := make([]int, 0, 3)
	for _, strDim := range strings.Split(lxwxh, "x") {
		intDim, err := strconv.Atoi(strDim)
		if err != nil {
			log.Fatal(err)
		}
		dimensions = append(dimensions, intDim)
	}
	if len(dimensions) != 3 {
		log.Fatal("Invalid dimensions for ", lxwxh)
	}
	return dimensions
}

/// dimensions - array of integers containing L, W, H
/// return area needed for wrapping
func calculateWrappingArea(dimensions []int) int {
	sort.Ints(dimensions)
	val1 := dimensions[0] * dimensions[1] // this is the minimum
	val2 := dimensions[0] * dimensions[2]
	val3 := dimensions[1] * dimensions[2]
	return 3*val1 + 2*val2 + 2*val3
}

/// dimensions - array of integers containing L, W, H
/// return length of ribbons
func calculateRibbonLength(dimensions []int) int {
	sort.Ints(dimensions)
	length := dimensions[0] * dimensions[1] * dimensions[2]
	return length + 2*(dimensions[0]+dimensions[1])
}
