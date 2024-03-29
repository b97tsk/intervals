package intervals

import "sort"

// Union returns the set of elements that are in either x, or y, or both.
func (x Set[E]) Union(y Set[E]) Set[E] {
	return Union(nil, x, y)
}

// Union returns the set of elements that are in either x, or y, or both,
// overwriting z. z must not be x or y and z must not be used after.
func Union[E Elem[E]](z, x, y Set[E]) Set[E] {
	z = z[:0]

	for {
		if len(x) == 0 || len(y) == 0 {
			if len(x) == 0 {
				x, y = y, x
			}

			return append(z, x...)
		}

		if x[0].High.Compare(y[0].High) > 0 {
			x, y = y, x
		}

		r := y[0]
		y = y[1:]

		i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.Low) >= 0 })
		z = append(z, x[:i]...)
		x = x[i:]

		if len(x) != 0 && x[0].Low.Compare(r.Low) < 0 {
			r.Low = x[0].Low
		}

	Again:
		j := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.High) > 0 })
		x = x[j:]

		if len(x) != 0 && x[0].Low.Compare(r.High) <= 0 {
			r.High = x[0].High
			x, y = y, x[1:]

			goto Again
		}

		z = append(z, r)
	}
}
