package intervals

import "sort"

// Intersection returns the set of elements that are in both x and y.
func (x Set[E]) Intersection(y Set[E]) Set[E] {
	return Intersection(nil, x, y)
}

// Intersection returns the set of elements that are in both x and y,
// overwriting z. z must not be x or y and z must not be used after.
func Intersection[E Elem[E]](z, x, y Set[E]) Set[E] {
	z = z[:0]

	for {
		if len(x) == 0 {
			x, y = y, x
		}

		if len(y) == 0 {
			return z
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
			start := len(z)
			z = append(z, x[:j]...)

			if r0 := &z[start]; r0.Low.Compare(r.Low) < 0 {
				r0.Low = r.Low
			}

			if r1 := &z[len(z)-1]; r1.High.Compare(r.High) > 0 {
				r1.High = r.High
				j--
			}

			x = x[j:]
		}
	}
}
