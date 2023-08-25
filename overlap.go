package intervals

import "sort"

// Overlaps reports whether the intersection of x and y is not empty.
func (x Set[E]) Overlaps(y Set[E]) bool {
	for {
		if len(x) < len(y) {
			x, y = y, x
		}

		if len(y) == 0 {
			return false
		}

		if x[0].High.Compare(y[0].High) > 0 {
			x, y = y, x
		}

		r := y[0]
		y = y[1:]

		i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.Low) > 0 })
		x = x[i:]
		j := sort.Search(len(x), func(i int) bool { return x[i].Low.Compare(r.High) >= 0 })

		if j > 0 {
			return true
		}
	}
}
