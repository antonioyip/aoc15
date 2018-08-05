package main

import (
	"fmt"
	"io/ioutil"
)

type Grid [102][102]int

func main() {
	inputs, err := ioutil.ReadFile("day18.input")
	if err != nil {
		panic(err)
	}

	grid := Grid{}
	row := 1
	col := 1
	for _, char := range inputs {
		switch char {
		case '#':
			grid[row][col] = 1
			row++
		case '.':
			grid[row][col] = 0
			row++
		case '\n':
			row = 1
			col++
		default:
			fmt.Println("Unexpected ", char)
		}
	}

	tempGrid := grid
	for i := 0; i < 100; i++ {
		tempGrid = flipLights(tempGrid)
	}
	fmt.Println(countOn(tempGrid))

	tempGrid = grid
	for i := 0; i < 100; i++ {
		tempGrid = stickCorners(flipLights(tempGrid))
	}
	fmt.Println(countOn(tempGrid))
}

func stickCorners(inGrid Grid) Grid {
	inGrid[1][100] = 1
	inGrid[100][1] = 1
	inGrid[100][100] = 1
	inGrid[1][1] = 1
	return inGrid
}

// apply rules for simultaneously flipping all lights
func flipLights(inGrid Grid) Grid {
	outGrid := Grid{}
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			on := neighborsOn(inGrid, i, j)
			if inGrid[i][j] == 1 {
				if on >= 2 && on <= 3 {
					outGrid[i][j] = 1
				} else {
					outGrid[i][j] = 0
				}
			} else { // inGrid[i][j] == 0
				if on == 3 {
					outGrid[i][j] = 1
				} else {
					outGrid[i][j] = 0
				}
			}
		}
	}

	return outGrid
}

// given grid[row][col], return the number of neighbors that are on
func neighborsOn(grid Grid, row int, col int) int {
	on := -grid[row][col]
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			on += grid[row+i][col+j]
		}
	}
	return on
}

func countOn(grid Grid) int {
	count := 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}
