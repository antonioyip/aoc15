package main

import (
	"aoc15/day21/rpg"
	"fmt"
)

func main() {
	test()

	// GetNextKit is a generator; abort is used to shutdown generator
	abort := make(chan struct{})

	// part 1
	lowestCost := 99999
	for kit := range rpg.GetNextKit(abort) {
		hero := *rpg.NewFighter(100, 0, 0)
		boss := *rpg.NewFighter(109, 8, 2)
		// equip hero
		currentCost := hero.Equip(kit)
		if battle(hero, boss) {
			// hero wins
			if currentCost < lowestCost {
				lowestCost = currentCost
			}
		}
	}
	fmt.Println(lowestCost)

	// part 2
	highestCost := 0
	for kit := range rpg.GetNextKit(abort) {
		hero := *rpg.NewFighter(100, 0, 0)
		boss := *rpg.NewFighter(109, 8, 2)
		// equip hero
		currentCost := hero.Equip(kit)
		if !battle(hero, boss) {
			// hero loses
			if highestCost < currentCost {
				highestCost = currentCost
			}
		}
	}
	fmt.Println(highestCost)
}

// return true if f1 wins
func battle(f1 rpg.Fighter, f2 rpg.Fighter) bool {
	for {
		f1.Attack(&f2)
		if f2.IsDead() {
			return true
		}
		f2.Attack(&f1)
		if f1.IsDead() {
			return false
		}
	}
}

func test() {
	hero := *rpg.NewFighter(8, 5, 5)
	boss := *rpg.NewFighter(12, 7, 2)
	fmt.Println(battle(hero, boss))
}
