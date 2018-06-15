package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// position of santa or robot
type carrier struct {
	row int
	col int
}

// move santa or robot
func (c *carrier) move(direction byte) {
	if direction == '^' {
		c.row++
	} else if direction == 'v' {
		c.row--
	} else if direction == '<' {
		c.col--
	} else if direction == '>' {
		c.col++
	} else {
		log.Fatal("Unexpected direction ", direction)
	}
}

func visitHouses(directions string) int {
	// houses[row][col] = presents
	houses := map[int]map[int]int{}

	// initial house
	santa := carrier{row: 0, col: 0}
	robot := carrier{row: 0, col: 0}
	houses[0] = map[int]int{}
	houses[0][0] = 2
	totalHouses := 1
	//fmt.Printf("%s is at %d %d\t%s is at %d %d\n", "santa", santa.row, santa.col, "robot", robot.row, robot.col)

	for i := 0; i < len(directions); i++ {
		agent := &santa
		if i%2 == 1 {
			agent = &robot
		}
		agent.move(directions[i])
		//fmt.Printf("%s is at %d %d\t%s is at %d %d\n", "santa", santa.row, santa.col, "robot", robot.row, robot.col)
		if houses[agent.row] == nil {
			houses[agent.row] = map[int]int{}
		}
		presents := houses[agent.row][agent.col] + 1
		houses[agent.row][agent.col] = presents
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

	test("^v", 3)
	test("^v^v", 5)
	test("^>v<", 3)
}
