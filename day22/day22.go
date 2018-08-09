package main

import (
	"aoc15/day22/rpg"
	"fmt"
)

func main() {
	hero := rpg.NewFighter(100, 500)
	boss := rpg.NewFighter(100, 0)
	//hero.AddEffect(rpg.NewShieldEffect())
	hero.AddAction(rpg.NewShield())
	hero.BattleAction(boss)
	for i := 0; i < 6; i++ {
		hero.ApplyEffects()
		hero.RemoveExpiredEffects()
	}
	boss.AddAction(rpg.NewMelee())
	boss.BattleAction(hero)
	fmt.Println(hero)
	fmt.Println(boss)
}
