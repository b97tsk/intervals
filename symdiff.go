package intervals

import "sort"

// SymmetricDifference returns the set of elements that are in one of x and y,
// but not in both.
func (x Set[E]) SymmetricDifference(y Set[E]) Set[E] {
	return SymmetricDifference(nil, x, y)
}

// SymmetricDifference returns the set of elements that are in one of x and y,
// but not in both, overwriting z. z must not be x or y and z must not be used
// after.
func SymmetricDifference[E Elem[E]](z, x, y Set[E]) Set[E] {
	z = z[:0]

	for {
		if len(x) == 0 {
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

func appendInterval[E Elem[E]](x Set[E], r Interval[E]) Set[E] {
	if n := len(x); n != 0 {
		if r1 := &x[n-1]; r1.High.Compare(r.Low) == 0 {
			r1.High = r.High
			return x
		}
	}

	return append(x, r)
}

func appendIntervals[E Elem[E]](x Set[E], s ...Interval[E]) Set[E] {
	if len(s) == 0 {
		return x
	}

	return append(appendInterval(x, s[0]), s[1:]...)
}
