package rpg

type Effect interface {
	Apply(f *Fighter)
	Expire(f *Fighter) bool
}

///////////////////////////////////

type ShieldEffect struct {
	duration int
	armor    int
}

func (e *ShieldEffect) Apply(f *Fighter) {
	if e.duration > 0 {
		f.armor = e.armor
		e.duration--
	}
}

func (e ShieldEffect) Expire(f *Fighter) bool {
	if e.duration < 1 {
		f.armor = 0
		return true
	}
	return false
}

///////////////////////////////////

type PoisonEffect struct {
	duration int
	damage   int
}

func (e *PoisonEffect) Apply(f *Fighter) {
	if e.duration > 0 {
		f.health -= e.damage
		e.duration--
	}
}

func (e PoisonEffect) Expire(f *Fighter) bool {
	if e.duration < 1 {
		return true
	}
	return false
}

///////////////////////////////////

type RechargeEffect struct {
	duration int
	mana     int
}

func (e *RechargeEffect) Apply(f *Fighter) {
	if e.duration > 0 {
		f.mana += e.mana
		e.duration--
	}
}

func (e RechargeEffect) Expire(f *Fighter) bool {
	if e.duration < 1 {
		return true
	}
	return false
}
