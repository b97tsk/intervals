package intervals

func combine[E Elem[E]](
	op func(x, y, out Set[E]) Set[E],
	sets ...Set[E],
) Set[E] {
	if len(sets) == 0 {
		return nil
	}

	var x, y Set[E]

	x = sets[0]
	xIsSets0 := true

	for _, set := range sets[1:] {
		y = op(x, set, y)

		if xIsSets0 {
			x = nil
			xIsSets0 = false
		}

		x, y = y, x
	}

	if xIsSets0 {
		// Always return a distinct set (unless it's nil).
		x = append(Set[E](nil), x...)
	}

	return x
}
