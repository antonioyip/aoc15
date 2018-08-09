package rpg

type Action interface {
	// returns mana used
	Execute(caster *Fighter, opponent *Fighter) int
	Usable(caster *Fighter, opponent *Fighter) bool
}

///////////////////////////////////

type Melee struct {
	cost   int
	damage int
}

func NewMelee() *Melee {
	return &Melee{0, 10}
}

func (m Melee) Usable(caster *Fighter, opponent *Fighter) bool {
	return true
}

func (m Melee) Execute(caster *Fighter, opponent *Fighter) int {
	calculatedDamage := m.damage - opponent.armor
	if calculatedDamage < 1 {
		calculatedDamage = 1
	}
	opponent.health -= calculatedDamage
	return m.cost
}

///////////////////////////////////

type Missile struct {
	cost   int
	damage int
}

func NewMissile() *Missile {
	return &Missile{53, 4}
}

func (m Missile) Usable(caster *Fighter, opponent *Fighter) bool {
	return caster.mana >= m.cost
}

func (m Missile) Execute(caster *Fighter, opponent *Fighter) int {
	opponent.health -= m.damage
	caster.mana -= m.cost
	return m.cost
}

///////////////////////////////////

type Drain struct {
	cost   int
	damage int
	heal   int
}

func NewDrain() *Drain {
	return &Drain{73, 2, 2}
}

func (d Drain) Usable(caster *Fighter, opponent *Fighter) bool {
	return caster.mana >= d.cost
}

func (d Drain) Execute(caster *Fighter, opponent *Fighter) int {
	opponent.health -= d.damage
	caster.health += d.heal
	caster.mana -= d.cost
	return d.cost
}

///////////////////////////////////

type Shield struct {
	cost     int
	duration int
	armor    int
}

func NewShield() *Shield {
	return &Shield{113, 6, 7}
}

func (s Shield) Usable(caster *Fighter, opponent *Fighter) bool {
	if caster.mana < s.cost {
		return false
	}
	// for _, e := range caster.effects {
	// 	if _, ok := e.(ShieldEffect); ok {
	// 		return false
	// 	}
	// }
	return true
}

func (s Shield) Execute(caster *Fighter, opponent *Fighter) int {
	caster.mana -= s.cost
	caster.AddEffect(&ShieldEffect{s.duration, s.armor})
	return s.cost
}

///////////////////////////////////

type Poison struct {
	cost     int
	duration int
	damage   int
}

func NewPoison() *Poison {
	return &Poison{173, 6, 3}
}

func (p Poison) Usable(caster *Fighter, opponent *Fighter) bool {
	if caster.mana < p.cost {
		return false
	}
	// for _, e := range opponent.effects {
	// 	if _, ok := e.(PoisonEffect); ok {
	// 		return false
	// 	}
	// }
	return true
}

func (p Poison) Execute(caster *Fighter, opponent *Fighter) int {
	caster.mana -= p.cost
	opponent.AddEffect(&PoisonEffect{p.duration, p.damage})
	return p.cost
}

///////////////////////////////////

type Recharge struct {
	cost     int
	duration int
	mana     int
}

func NewRecharge() *Recharge {
	return &Recharge{229, 5, 101}
}

func (r Recharge) Usable(caster *Fighter, opponent *Fighter) bool {
	if caster.mana < r.cost {
		return false
	}
	// for _, e := range caster.effects {
	// 	if _, ok := e.(RechargeEffect); ok {
	// 		return false
	// 	}
	// }
	return true
}

func (r Recharge) Execute(caster *Fighter, opponent *Fighter) int {
	caster.mana -= r.cost
	caster.AddEffect(&RechargeEffect{r.duration, r.mana})
	return r.cost
}
