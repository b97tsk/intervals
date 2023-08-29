// Package intervals is a library for manipulating sets of intervals.
package intervals

import "sort"

// Elem is the type set containing all supported element types.
type Elem[E any] interface {
	Compare(E) int
	Next() E
}

// An Interval is a half-open continuous range of elements.
type Interval[E Elem[E]] struct {
	Low  E // inclusive
	High E // exclusive
}

// Equal reports whether r is identical to r2.
func (r Interval[E]) Equal(r2 Interval[E]) bool {
	return r.Low.Compare(r2.Low) == 0 && r.High.Compare(r2.High) == 0
}

// Unwrap returns r.Low, r.High.
func (r Interval[E]) Unwrap() (E, E) {
	return r.Low, r.High
}

// A Set is a slice of separate intervals sorted in ascending order.
// The zero value for a Set, i.e. a nil Set, is an empty set.
//
// Since Interval is half-open, you can never add the maximum value of E into
// a Set.
type Set[E Elem[E]] []Interval[E]

// Add adds a single element into x.
func (x *Set[E]) Add(v E) {
	x.AddRange(v, v.Next())
}

// AddRange adds range [lo, hi) into x.
func (x *Set[E]) AddRange(lo, hi E) {
	s := *x

	i := sort.Search(len(s), func(i int) bool { return s[i].Low.Compare(lo) > 0 })
	j := sort.Search(len(s), func(i int) bool { return s[i].High.Compare(hi) > 0 })

	// ┌────────┬─────────────────────────────────────────┐
	// │        │    j-1        j        i-1        i     │
	// │ Case 1 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │        |<- hi  ->|   |<- lo  ->|        │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    j-1        j         i               │
	// │ Case 2 │  |-----|   |-----|   |-----|            │
	// │        │            |<- lo  ->|                  │
	// │        │        |<- hi  ->|                      │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1       i,j                        │
	// │ Case 3 │  |-----|   |-----|                      │
	// │        │  |<- lo  ->|                            │
	// │        │        |<- hi  ->|                      │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i         j               │
	// │ Case 4 │  |-----|   |-----|   |-----|            │
	// │        │  |<- lo  ->|     |<- hi  ->|            │
	// │        │                                         │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i        j-1        j     │
	// │ Case 5 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │  |<- lo  ->|               |<- hi  ->|  │
	// └────────┴─────────────────────────────────────────┘

	if i > j { // Case 1 and 2.
		return
	}

	// Case 3, 4 and 5.

	if i > 0 {
		if r := &s[i-1]; r.High.Compare(lo) >= 0 {
			lo = r.Low
			i--
		}
	}

	if j < len(s) {
		if r := &s[j]; r.Low.Compare(hi) <= 0 {
			hi = r.High
			j++
		}
	}

	if i == j { // Case 3 (where lo and hi overlap with each other).
		if lo.Compare(hi) < 0 {
			s = append(s, Interval[E]{})
			copy(s[i+1:], s[i:])
			s[i] = Interval[E]{lo, hi}
			*x = s
		}

		return
	}

	s[i] = Interval[E]{lo, hi}
	s = append(s[:i+1], s[j:]...)
	*x = s
}

// Delete removes a single element from x.
func (x *Set[E]) Delete(v E) {
	x.DeleteRange(v, v.Next())
}

// DeleteRange removes range [lo, hi) from x.
func (x *Set[E]) DeleteRange(lo, hi E) {
	s := *x

	i := sort.Search(len(s), func(i int) bool { return s[i].High.Compare(lo) > 0 })
	// j := sort.Search(len(s), func(i int) bool { return s[i].Low.Compare(hi) > 0 })

	// ┌────────┬─────────────────────────────────────────┐
	// │        │    j-1        j        i-1        i     │
	// │ Case 1 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │  |<- hi  ->|               |<- lo  ->|  │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    j-1        j         i               │
	// │ Case 2 │  |-----|   |-----|   |-----|            │
	// │        │  |<- hi  ->|     |<- lo  ->|            │
	// │        │                                         │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1       i,j                        │
	// │ Case 3 │  |-----|   |-----|                      │
	// │        │        |<- lo  ->|                      │
	// │        │  |<- hi  ->|                            │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i         j               │
	// │ Case 4 │  |-----|   |-----|   |-----|            │
	// │        │        |<- lo  ->|                      │
	// │        │            |<- hi  ->|                  │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i        j-1        j     │
	// │ Case 5 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │        |<- lo  ->|   |<- hi  ->|        │
	// └────────┴─────────────────────────────────────────┘

	// Optimized, j >= i.
	t := s[i:]
	j := i + sort.Search(len(t), func(i int) bool { return t[i].Low.Compare(hi) > 0 })

	if i == j { // Case 1, 2 and 3.
		return
	}

	if i == j-1 { // Case 4.
		if r := &s[i]; r.Low.Compare(lo) < 0 {
			if r.High.Compare(hi) > 0 {
				if lo.Compare(hi) < 0 {
					s = append(s, Interval[E]{})
					copy(s[j:], s[i:])
					s[i].High = lo
					s[j].Low = hi
					*x = s
				}
			} else {
				r.High = lo
			}
		} else {
			if r.High.Compare(hi) > 0 {
				r.Low = hi
			} else {
				s = append(s[:i], s[j:]...)
				*x = s
			}
		}

		return
	}

	// Case 5.

	if r := &s[i]; r.Low.Compare(lo) < 0 {
		r.High = lo
		i++
	}

	if r := &s[j-1]; r.High.Compare(hi) > 0 {
		r.Low = hi
		j--
	}

	s = append(s[:i], s[j:]...)
	*x = s
}

// Contains reports whether x contains a single element.
func (x Set[E]) Contains(v E) bool {
	i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(v) > 0 })
	return i < len(x) && x[i].Low.Compare(v) <= 0
}

// ContainsRange reports whether x contains every element in range [lo, hi).
func (x Set[E]) ContainsRange(lo, hi E) bool {
	i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(lo) > 0 })
	return i < len(x) && x[i].Low.Compare(lo) <= 0 && x[i].High.Compare(hi) >= 0 && lo.Compare(hi) < 0
}

// Equal reports whether x is identical to y.
func (x Set[E]) Equal(y Set[E]) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if !x[i].Equal(y[i]) {
			return false
		}
	}

	return true
}

// Extent returns the smallest Interval that contains every element in x.
//
// If x is empty, Extent returns the zero value.
func (x Set[E]) Extent() Interval[E] {
	if len(x) == 0 {
		return Interval[E]{}
	}

	return Interval[E]{
		Low:  x[0].Low,
		High: x[len(x)-1].High,
	}
}

// IsSubsetOf reports whether elements in x are all in y.
func (x Set[E]) IsSubsetOf(y Set[E]) bool {
	for i := range x {
		if !y.ContainsRange(x[i].Unwrap()) {
			return false
		}
	}

	return true
}
