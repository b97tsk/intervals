// Package intervals is a library for manipulating sets of intervals.
package intervals

import (
	"slices"
	"sort"
)

// Elem is the type set containing all supported element types.
type Elem[E any] interface {
	Compare(E) int
}

// An Enum is a computably enumerable Elem. There exists a computable
// enumeration that lists elements, in a valid Interval of Enum, in an
// increasing ordering determined by the Compare method.
type Enum[E any] interface {
	Elem[E]
	Next() E
}

// An Interval is a half-open continuous range of elements.
type Interval[E Elem[E]] struct {
	Low  E // inclusive
	High E // exclusive
}

// Unit returns an Interval that only contains a single element v.
// If there is no element next to v, Unit returns an invalid Interval.
// Note that calling Set method on an invalid Interval panics.
func Unit[E Enum[E]](v E) Interval[E] {
	return Range(v, v.Next())
}

// Range returns an Interval of range [lo, hi).
// If lo.Compare(hi) >= 0, Range returns an invalid Interval.
// Note that calling Set method on an invalid Interval panics.
func Range[E Elem[E]](lo, hi E) Interval[E] {
	return Interval[E]{lo, hi}
}

// Equal reports whether r is identical to r2.
func (r Interval[E]) Equal(r2 Interval[E]) bool {
	return r.Low.Compare(r2.Low) == 0 && r.High.Compare(r2.High) == 0
}

// IsValid reports whether r is a valid Interval.
// A valid Interval r either satisfies r.Low.Compare(r.High) < 0 or equals to
// the zero value.
func (r Interval[E]) IsValid() bool {
	return r.test() <= 0
}

// Set returns the set of elements that are in r.
// If r is the zero value, Set returns an empty set.
// If r is an invalid Interval, Set panics.
func (r Interval[E]) Set() Set[E] {
	switch r.test() {
	case -1:
		return Set[E]{r}
	case 0:
		return nil
	}

	panic("invalid Interval")
}

func (r Interval[E]) test() int {
	switch c := r.Low.Compare(r.High); {
	case c < 0:
		return -1
	case c == 0 && r.Equal(Interval[E]{}):
		return 0
	default:
		return +1
	}
}

// A Set is a slice of separate intervals sorted in ascending order.
// The zero value for a Set, i.e. a nil Set, is an empty set.
//
// Since Intervals are half-open, the maximum value of E cannot be added into
// a Set.
type Set[E Elem[E]] []Interval[E]

// Collect returns the set of elements that are in any of s.
//
// Collect performs better if s are sorted in ascending order by the Low field.
func Collect[E Elem[E]](s ...Interval[E]) Set[E] {
	return CollectInto(nil, s...)
}

// CollectInto returns the set of elements that are in any of s, overwriting x.
// x and s can be the same slice. x must not be used after.
//
// CollectInto performs better if s are sorted in ascending order by the Low
// field.
func CollectInto[E Elem[E]](x Set[E], s ...Interval[E]) Set[E] {
	x = x[:0]

	for _, r := range s {
		if len(x) == 0 {
			if r.Low.Compare(r.High) < 0 {
				x = append(x, r)
			}

			continue
		}

		switch r1 := &x[len(x)-1]; {
		case r1.High.Compare(r.Low) < 0:
			if r.Low.Compare(r.High) < 0 {
				x = append(x, r)
			}
		case r1.Low.Compare(r.Low) <= 0:
			if r1.High.Compare(r.High) < 0 {
				r1.High = r.High
			}
		default:
			x = addRange(x, r.Low, r.High)
		}
	}

	return x
}

// Add adds zero or more Intervals into x, returning the modified Set.
func Add[E Elem[E]](x Set[E], s ...Interval[E]) Set[E] {
	for _, r := range s {
		x = addRange(x, r.Low, r.High)
	}

	return x
}

// addRange adds range [lo, hi) into x, returning the modified Set.
func addRange[E Elem[E]](x Set[E], lo, hi E) Set[E] {
	i := sort.Search(len(x), func(i int) bool { return x[i].Low.Compare(lo) > 0 })
	// j := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(hi) > 0 })

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

	// Optimized: j >= i-1.
	off := max(0, i-1)
	z := x[off:]
	j := off + sort.Search(len(z), func(i int) bool { return z[i].High.Compare(hi) > 0 })

	if i > j { // Case 1 and 2.
		return x
	}

	// Case 3, 4 and 5.

	if i > 0 {
		if r := &x[i-1]; r.High.Compare(lo) >= 0 {
			lo = r.Low
			i--
		}
	}

	if j < len(x) {
		if r := &x[j]; r.Low.Compare(hi) <= 0 {
			hi = r.High
			j++
		}
	}

	if i == j { // Case 3 (where lo and hi overlap).
		if lo.Compare(hi) >= 0 { // Get rid of the devil first.
			return x
		}
	}

	return slices.Replace(x, i, j, Range(lo, hi))
}

// Delete removes zero or more Intervals from x, returning the modified Set.
func Delete[E Elem[E]](x Set[E], s ...Interval[E]) Set[E] {
	for _, r := range s {
		x = deleteRange(x, r.Low, r.High)
	}

	return x
}

// deleteRange removes range [lo, hi) from x, returning the modified Set.
func deleteRange[E Elem[E]](x Set[E], lo, hi E) Set[E] {
	i := sort.Search(len(x), func(i int) bool { return x[i].High.Compare(lo) > 0 })
	// j := sort.Search(len(x), func(i int) bool { return x[i].Low.Compare(hi) > 0 })

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

	// Optimized: j >= i.
	z := x[i:]
	j := i + sort.Search(len(z), func(i int) bool { return z[i].Low.Compare(hi) > 0 })

	if i == j { // Case 1, 2 and 3.
		return x
	}

	if i == j-1 { // Case 4.
		if lo.Compare(hi) >= 0 { // Get rid of the devil first.
			return x
		}
	}

	// Case 4 and 5.

	v := make([]Interval[E], 0, 2)

	if r := &x[i]; r.Low.Compare(lo) < 0 {
		v = append(v, Range(r.Low, lo))
	}

	if r := &x[j-1]; r.High.Compare(hi) > 0 {
		v = append(v, Range(hi, r.High))
	}

	return slices.Replace(x, i, j, v...)
}

// Contains reports whether x contains every element in range [r.Low, r.High).
func (x Set[E]) Contains(r Interval[E]) bool {
	x = x[sort.Search(len(x), func(i int) bool { return x[i].High.Compare(r.Low) > 0 }):]
	return len(x) != 0 && x[0].Low.Compare(r.Low) <= 0 && x[0].High.Compare(r.High) >= 0 && r.Low.Compare(r.High) < 0
}

// ContainsUnit reports whether x contains a single element v.
// Faster than x.Contains(Unit(v)).
func (x Set[E]) ContainsUnit(v E) bool {
	x = x[sort.Search(len(x), func(i int) bool { return x[i].High.Compare(v) > 0 }):]
	return len(x) != 0 && x[0].Low.Compare(v) <= 0
}

// Equal reports whether x is identical to y.
func (x Set[E]) Equal(y Set[E]) bool {
	return slices.EqualFunc(x, y, Interval[E].Equal)
}

// Extent returns the smallest Interval that contains every element in x.
//
// If x is empty, Extent returns the zero value.
func (x Set[E]) Extent() Interval[E] {
	if len(x) == 0 {
		return Interval[E]{}
	}

	return Range(x[0].Low, x[len(x)-1].High)
}
