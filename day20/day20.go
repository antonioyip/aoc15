package main

import (
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	goal := 34000000
	houses := make([]int, 1000000)
	for i := 1; i < len(houses); i++ {
		for j := i; j < len(houses); j += i {
			houses[j] += i * 10
			if houses[j] > goal {
				fmt.Println(j)
				return
			}
		}
	}
}

func part2() {
	goal := 34000000
	houses := make([]int, 1000000)
	for i := 1; i < len(houses); i++ {
		for j := i; j < i*50 && j < len(houses); j += i {
			houses[j] += i * 11
			if houses[j] > goal {
				fmt.Println(j)
				return
			}
		}
	}
}
