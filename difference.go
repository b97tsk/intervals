package intervals

import "sort"

// Difference returns the set of elements that are in x, but not in y.
func (x Set[E]) Difference(y Set[E]) Set[E] {
	return Difference(nil, x, y)
}

// Difference returns the set of elements that are in x, but not in y,
// overwriting z. z must not be x or y and z must not be used after.
func Difference[E Elem[E]](z, x, y Set[E]) Set[E] {
	z = z[:0]

	inv := false

	for {
		if len(x) < len(y) {
			x, y = y, x
			inv = !inv
		}

		if len(y) == 0 {
			if !inv {
				z = append(z, x...)
			}

			return z
		}

		if x[0].High.Compare(y[0].High) > 0 {
			x, y = y, x
			inv = !inv
		}

		r := y[0]
		y = y[1:]

		i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.Low) > 0 })

		if !inv {
			z = append(z, x[:i]...)
		}

		x = x[i:]

	Again:
		j := sort.Search(len(x), func(i int) bool { return x[i].Low.Compare(r.High) >= 0 })

		if j == 0 {
			if inv {
				z = append(z, r)
			}

			continue
		}

		lo := x[0].Low

		switch c := lo.Compare(r.Low); {
		case c < 0:
			if !inv {
				z = append(z, Range(lo, r.Low))
			}
		case c > 0:
			if inv {
				z = append(z, Range(r.Low, lo))
			}
		}

		if inv {
			for i := 0; i < j-1; i++ {
				z = append(z, Range(x[i].High, x[i+1].Low))
			}
		}

		hi := x[j-1].High
		x = x[j:]

		switch c := hi.Compare(r.High); {
		case c < 0:
			if inv {
				z = append(z, Range(hi, r.High))
			}
		case c > 0:
			r.Low, r.High = r.High, hi
			x, y = y, x
			inv = !inv

			goto Again
		}
	}
}
