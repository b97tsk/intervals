package intervals

import "sort"

// Union returns the set of elements that are in either x, or y, or both.
func (x Set[E]) Union(y Set[E]) Set[E] {
	return UnionInto(nil, x, y)
}

// Union returns the set of elements that are in any of sets.
func Union[E Elem[E]](sets ...Set[E]) Set[E] {
	return combine(UnionInto, sets...)
}

// UnionInto returns the set of elements that are in either x, or y, or both,
// overwriting z. z must not be x or y; z must not be used after.
func UnionInto[E Elem[E]](z, x, y Set[E]) Set[E] {
	z = z[:0]

	for {
		if len(x) < len(y) {
			x, y = y, x
		}

		if len(y) == 0 {
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
