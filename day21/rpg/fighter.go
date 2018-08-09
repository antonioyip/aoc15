package rpg

type Fighter struct {
	hitpoints int
	damage    int
	armor     int
}

func NewFighter(hitpoints, damage, armor int) *Fighter {
	return &Fighter{hitpoints, damage, armor}
}

func (f1 Fighter) Attack(f2 *Fighter) {
	calculatedDamage := f1.damage - f2.armor
	if calculatedDamage < 1 {
		calculatedDamage = 1
	}
	f2.hitpoints -= calculatedDamage
	//fmt.Println(calculatedDamage, f2.hitpoints)
}

func (f *Fighter) Equip(kit []Gear) int {
	cost := 0
	for _, g := range kit {
		cost += g.cost
		f.armor += g.armor
		f.damage += g.damage
	}
	return cost
}

func (f Fighter) IsDead() bool {
	return f.hitpoints <= 0
}
