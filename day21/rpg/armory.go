package rpg

type Gear struct {
	cost   int
	damage int
	armor  int
}

// 1 weapon
// 0-1 armor
// 0-2 rings
func GetNextKit(abort <-chan struct{}) <-chan []Gear {

	ch := make(chan []Gear)

	go func() {
		defer close(ch)
		weapons := []Gear{
			{8, 4, 0},
			{10, 5, 0},
			{25, 6, 0},
			{40, 7, 0},
			{74, 8, 0},
		}

		armor := []Gear{
			{0, 0, 0}, // no armor
			{13, 0, 1},
			{31, 0, 2},
			{53, 0, 3},
			{75, 0, 4},
			{102, 0, 5},
		}

		rings := []Gear{
			{0, 0, 0}, // 1 ring
			{25, 1, 0},
			{50, 2, 0},
			{100, 3, 0},
			{20, 0, 1},
			{40, 0, 2},
			{80, 0, 3},
		}

		// handle the no ring use case
		noRings := []Gear{weapons[0], armor[0], rings[0], rings[0]}
		select {
		case ch <- noRings:
		case <-abort:
			return
		}

		for _, w := range weapons {
			for _, a := range armor {
				for _, r1 := range rings {
					for _, r2 := range rings {
						if r1 == r2 {
							// no duplicate rings
							continue
						}
						select {
						case ch <- []Gear{w, a, r1, r2}:
						case <-abort:
							return
						}
					}
				}
			}
		}
	}()

	return ch
}
