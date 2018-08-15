package main

import (
	"fmt"
)

type State struct {
	heroHealth       int
	heroMana         int
	heroArmor        int
	bossHealth       int
	bossDamage       int
	shieldDuration   int
	poisonDuration   int
	rechargeDuration int
	manaUsed         int
	minMana          int
}

type Spell struct {
	name           string
	cost           int
	damage         int
	heal           int
	updateDuration func(state *State)
}

func (state State) IsVictory() bool {
	return state.bossHealth < 1
}

func (state State) IsDefeat() bool {
	return state.heroHealth < 1
}

func (spell Spell) Available(state State) bool {
	// check sufficient mana
	if spell.cost > state.heroMana {
		return false
	}
	// check if active
	switch spell.name {
	case "Shield":
		if state.shieldDuration > 0 {
			return false
		}
	case "Poison":
		if state.poisonDuration > 0 {
			return false
		}
	case "Recharge":
		if state.rechargeDuration > 0 {
			return false
		}
	}
	return true
}

func (spell Spell) Cast(state *State) {
	state.heroMana -= spell.cost
	state.heroHealth += spell.heal
	state.bossHealth -= spell.damage
	spell.updateDuration(state)
	state.manaUsed += spell.cost
}

func ApplyEffects(state *State) {
	// apply
	if state.shieldDuration > 0 {
		state.heroArmor = 7
	}
	if state.poisonDuration > 0 {
		state.bossHealth -= 3
	}
	if state.rechargeDuration > 0 {
		state.heroMana += 101
	}
	// decay
	state.shieldDuration--
	state.poisonDuration--
	state.rechargeDuration--
	// expire
	if state.shieldDuration < 1 {
		state.heroArmor = 0
	}
}

func BossAttack(state *State) {
	damage := state.bossDamage - state.heroArmor
	if damage < 1 {
		damage = 1
	}
	state.heroHealth -= damage
}

var spells = []Spell{
	Spell{"Missile", 53, 4, 0, func(state *State) {}},
	Spell{"Drain", 73, 2, 2, func(state *State) {}},
	Spell{"Shield", 113, 0, 0, func(state *State) { state.shieldDuration = 6 }},
	Spell{"Poison", 173, 0, 0, func(state *State) { state.poisonDuration = 6 }},
	Spell{"Recharge", 229, 0, 0, func(state *State) { state.rechargeDuration = 5 }},
}

func main() {
	state := State{50, 500, 0, 55, 8, 0, 0, 0, 0, 999999}
	fmt.Println(battle(state))
}

func battle(state State) int {
	for _, spell := range spells {
		if spell.Available(state) {
			success, manaUsed := heroTurn(spell, state)
			if success {
				if state.minMana > manaUsed {
					state.minMana = manaUsed
				}
			}
		}
	}
	return state.minMana
}

func heroTurn(spell Spell, state State) (bool, int) {
	if state.IsDefeat() {
		return false, 0
	}
	ApplyEffects(&state)
	if state.IsVictory() {
		return true, state.manaUsed
	}
	spell.Cast(&state)
	if state.manaUsed > state.minMana {
		// no longer most efficient
		return false, 0
	}
	if state.IsVictory() {
		return true, state.manaUsed
	}
	return bossTurn(state)
}

func bossTurn(state State) (bool, int) {
	ApplyEffects(&state)
	if state.IsVictory() {
		return true, state.manaUsed
	}

	BossAttack(&state)
	if state.IsDefeat() {
		return false, 0
	}

	victory := false
	for _, spell := range spells {
		if spell.Available(state) {
			success, manaUsed := heroTurn(spell, state)
			if success {
				victory = true
				if state.minMana > manaUsed {
					state.minMana = manaUsed
				}
			}
		}
	}

	if victory {
		return true, state.minMana
	}
	return false, 0
}
