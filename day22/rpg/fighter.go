package rpg

type Fighter struct {
	health  int
	mana    int
	armor   int
	effects []Effect
	actions []Action
}

func NewFighter(health int, mana int) *Fighter {
	return &Fighter{health: health, mana: mana}
}

func (f *Fighter) AddAction(a Action) {
	f.actions = append(f.actions, a)
}

func (f *Fighter) AddEffect(e Effect) {
	f.effects = append(f.effects, e)
}

func (f *Fighter) ApplyEffects() {
	for i := 0; i < len(f.effects); i++ {
		f.effects[i].Apply(f)
	}
}

func (f *Fighter) RemoveExpiredEffects() {
	tmp := f.effects[:0]
	for _, effect := range f.effects {
		if !effect.Expire(f) {
			tmp = append(tmp, effect)
		}
	}
	f.effects = tmp
}

func (f *Fighter) BattleAction(opponent *Fighter) {
	f.actions[0].Execute(f, opponent)
}
