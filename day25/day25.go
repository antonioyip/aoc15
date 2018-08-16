package main

import "fmt"

func main() {
	grid := makeSquareGrid(6)

	row, col := 0, 0
	for i := 1; i < 22; i++ {
		grid[row][col] = i
		row, col = next(row, col)
	}

	for i := range grid {
		fmt.Println(grid[i])
	}

	grid = makeSquareGrid(6000)
	lastValue := 20151125
	grid[0][0] = lastValue
	row, col = 0, 0
	for {
		if grid[2947-1][3029-1] != 0 {
			fmt.Println(grid[2947-1][3029-1])
			break
		}
		row, col = next(row, col)
		lastValue = (lastValue * 252533) % 33554393
		grid[row][col] = lastValue
	}
}

func makeSquareGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}

func next(row, col int) (int, int) {
	if row == 0 {
		return col + 1, 0
	}
	return row - 1, col + 1
}
