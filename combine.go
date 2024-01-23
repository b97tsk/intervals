package intervals

// Combine applies set operation op over sets, returning the end result.
// op can be any of [Difference], [Intersection], [SymmetricDifference] and
// [Union].
func Combine[E Elem[E]](op func(z, x, y Set[E]) Set[E], sets ...Set[E]) Set[E] {
	if len(sets) == 0 {
		return nil
	}

	var x, z Set[E]

	x = sets[0]
	xIsSets0 := true

	for _, y := range sets[1:] {
		z = op(z, x, y)

		if xIsSets0 {
			x = nil
			xIsSets0 = false
		}

		x, z = z, x
	}

	if xIsSets0 {
		// Always return a distinct set (unless it's nil).
		x = append(Set[E](nil), x...)
	}

	return x
}
