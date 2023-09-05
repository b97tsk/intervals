package intervals

import "sort"

// Difference returns the set of elements that are in x, but not in y.
func (x Set[E]) Difference(y Set[E]) Set[E] {
	return difference(x, y, nil)
}

// Difference returns the set of elements that are in sets[0], but not in
// any of sets[1:]. If sets is empty, Difference returns an empty set.
func Difference[E Elem[E]](sets ...Set[E]) Set[E] {
	return combine(difference, sets...)
}

func difference[E Elem[E]](x, y, out Set[E]) Set[E] {
	z := out[:0]

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

		switch lo := x[0].Low; lo.Compare(r.Low) {
		case -1:
			if !inv {
				z = append(z, Interval[E]{lo, r.Low})
			}
		case +1:
			if inv {
				z = append(z, Interval[E]{r.Low, lo})
			}
		}

		if inv {
			for i := 0; i < j-1; i++ {
				z = append(z, Interval[E]{x[i].High, x[i+1].Low})
			}
		}

		hi := x[j-1].High
		x = x[j:]

		switch hi.Compare(r.High) {
		case -1:
			if inv {
				z = append(z, Interval[E]{hi, r.High})
			}
		case +1:
			r.Low, r.High = r.High, hi
			x, y = y, x
			inv = !inv

			goto Again
		}
	}
}
