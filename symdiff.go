package intervals

import "sort"

// SymmetricDifference returns the set of elements that are in one of x and y,
// but not in both.
func (x Set[E]) SymmetricDifference(y Set[E]) Set[E] {
	return symmetricDifferenceInto(nil, x, y)
}

// SymmetricDifference returns the set of elements that are in an odd number
// of sets.
func SymmetricDifference[E Elem[E]](sets ...Set[E]) Set[E] {
	return combine(symmetricDifferenceInto, sets...)
}

// symmetricDifferenceInto returns the set of elements that are in one of x and
// y, but not in both, overwriting z. z must not be x or y; z must not be used
// after.
func symmetricDifferenceInto[E Elem[E]](z, x, y Set[E]) Set[E] {
	z = z[:0]

	for {
		if len(x) < len(y) {
			x, y = y, x
		}

		if len(y) == 0 {
			z = appendIntervals(z, x...)
			return z
		}

		if x[0].High.Compare(y[0].High) > 0 {
			x, y = y, x
		}

		r := y[0]
		y = y[1:]

		i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.Low) > 0 })
		z = appendIntervals(z, x[:i]...)
		x = x[i:]

	Again:
		j := sort.Search(len(x), func(i int) bool { return x[i].Low.Compare(r.High) >= 0 })

		if j == 0 {
			z = appendInterval(z, r)
			continue
		}

		lo := x[0].Low

		switch c := lo.Compare(r.Low); {
		case c < 0:
			z = appendInterval(z, Range(lo, r.Low))
		case c > 0:
			z = appendInterval(z, Range(r.Low, lo))
		}

		for i := 0; i < j-1; i++ {
			z = append(z, Range(x[i].High, x[i+1].Low))
		}

		hi := x[j-1].High
		x = x[j:]

		switch c := hi.Compare(r.High); {
		case c < 0:
			z = append(z, Range(hi, r.High))
		case c > 0:
			r.Low, r.High = r.High, hi
			x, y = y, x

			goto Again
		}
	}
}

func appendInterval[E Elem[E]](s Set[E], r Interval[E]) Set[E] {
	if n := len(s); n != 0 {
		if r1 := &s[n-1]; r1.High.Compare(r.Low) == 0 {
			r1.High = r.High
			return s
		}
	}

	return append(s, r)
}

func appendIntervals[E Elem[E]](s Set[E], rs ...Interval[E]) Set[E] {
	if len(rs) == 0 {
		return s
	}

	return append(appendInterval(s, rs[0]), rs[1:]...)
}
