package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func visitHouses(directions string) int {
	// houses[row][col] = presents
	houses := map[int]map[int]int{}

	// initial house
	row := 0
	col := 0
	totalHouses := 1
	houses[row] = map[int]int{}
	houses[row][col] = 1

	for _, direction := range directions {
		if direction == '^' {
			row++
		} else if direction == 'v' {
			row--
		} else if direction == '<' {
			col--
		} else if direction == '>' {
			col++
		} else {
			log.Fatal("Unexpected direction ", direction)
		}

		if houses[row] == nil {
			houses[row] = map[int]int{}
		}

		presents := houses[row][col] + 1
		houses[row][col] = presents
		if presents == 1 {
			totalHouses++
		}
	}
	return totalHouses
}

func test(input string, count int) {
	if visitHouses(input) == count {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func main() {
	input, err := ioutil.ReadFile("day3.input")
	if err != nil {
		log.Fatal(err)
	}
	directions := string(input)
	fmt.Printf("Houses visited: %d\n", visitHouses(directions))

	test("^^vv", 3)
	test("^v^v", 2)
	test("<>", 2)
	test(">", 2)
	test("<<>>", 3)
	test("^>", 3)
	test("^>v<", 4)
}
