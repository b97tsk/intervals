package intervals

import "sort"

// IsSubsetOf reports whether elements in x are all in y.
func (x Set[E]) IsSubsetOf(y Set[E]) bool {
	inv := false

	for {
		if len(x) < len(y) {
			x, y = y, x
			inv = !inv
		}

		if len(y) == 0 {
			return inv || len(x) == 0
		}

		if x[0].High.Compare(y[0].High) > 0 {
			x, y = y, x
			inv = !inv
		}

		r := y[0]
		y = y[1:]

		if inv {
			i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.Low) > 0 })
			x = x[i:]

			if len(x) == 0 || x[0].Low.Compare(r.Low) > 0 || x[0].High.Compare(r.High) < 0 {
				return false
			}

			continue
		}

		if x[0].Low.Compare(r.Low) < 0 {
			return false
		}

		j := sort.Search(len(x), func(i int) bool { return x[i].Low.Compare(r.High) >= 0 })

		if j > 0 {
			if x[j-1].High.Compare(r.High) > 0 {
				return false
			}

			x = x[j:]
		}
	}
}
