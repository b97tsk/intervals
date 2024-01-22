package intervals_test

import (
	"crypto/rand"
	"slices"
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func FuzzAdd(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		z := append(Set[elems.Uint8](nil), x...)

		for _, r := range y {
			z.Add(r)
		}

		if w := plainUnion(x, y); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x ∪ y = %v", w)
			t.Logf("x ∪ y = %v (actual)", z)
			t.Fail()
		}
	})
}

func FuzzDelete(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		z := append(Set[elems.Uint8](nil), x...)

		for _, r := range y {
			z.Delete(r)
		}

		if w := plainDifference(x, y); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x \\ y = %v", w)
			t.Logf("x \\ y = %v (actual)", z)
			t.Fail()
		}
	})
}

func FuzzContains(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		yes := true

		for _, r := range x {
			if !y.Contains(r) {
				yes = false
				break
			}
		}

		if yes != plainIsSubsetOf(x, y) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x is subset of y = %v", !yes)
			t.Logf("x is subset of y = %v (actual)", yes)
			t.Fail()
		}
	})
}

func FuzzIsSubsetOf(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		if yes := x.IsSubsetOf(y); yes != plainIsSubsetOf(x, y) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x is subset of y = %v", !yes)
			t.Logf("x is subset of y = %v (actual)", yes)
			t.Fail()
		}
	})
}

func FuzzIntersection(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		if z, w := x.Intersection(y), plainIntersection(x, y); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x ∩ y = %v", w)
			t.Logf("x ∩ y = %v (actual)", z)
			t.Fail()
		}
	})
}

func FuzzUnion(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		if z, w := x.Union(y), plainUnion(x, y); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x ∪ y = %v", w)
			t.Logf("x ∪ y = %v (actual)", z)
			t.Fail()
		}
	})
}

func FuzzDifference(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		if z, w := x.Difference(y), plainDifference(x, y); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x \\ y = %v", w)
			t.Logf("x \\ y = %v (actual)", z)
			t.Fail()
		}
	})
}

func FuzzSymmetricDifference(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		if z, w := x.SymmetricDifference(y), plainSymmetricDifference(x, y); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("y = %v", y)
			t.Logf("x △ y = %v", w)
			t.Logf("x △ y = %v (actual)", z)
			t.Fail()
		}
	})
}

func addRandomSeed(f *testing.F, n int) {
	data := make([]byte, n)
	args := make([]any, n)

	for i := 0; i < 10; i++ {
		if _, err := rand.Read(data); err != nil {
			f.Fatal(err)
		}

		for i, v := range data {
			args[i] = v
		}

		f.Add(args...)
	}
}

func fuzz(f *testing.F, ff func(t *testing.T, x, y Set[elems.Uint8])) {
	addRandomSeed(f, 32)

	f.Fuzz(func(
		t *testing.T,
		x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15,
		y0, y1, y2, y3, y4, y5, y6, y7, y8, y9, y10, y11, y12, y13, y14, y15 byte,
	) {
		xs := []byte{x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15}
		ys := []byte{y0, y1, y2, y3, y4, y5, y6, y7, y8, y9, y10, y11, y12, y13, y14, y15}

		slices.Sort(xs)
		slices.Sort(ys)

		var x, y []Interval[elems.Uint8]

		for i, j := 0, len(xs); i < j; i += 2 {
			if lo, hi := elems.Uint8(xs[i]), elems.Uint8(xs[i+1]); lo < hi {
				x = append(x, Range(lo, hi))
			}
		}

		for i, j := 0, len(ys); i < j; i += 2 {
			if lo, hi := elems.Uint8(ys[i]), elems.Uint8(ys[i+1]); lo < hi {
				y = append(y, Range(lo, hi))
			}
		}

		ff(t, plainUnion(x), plainUnion(y))
	})
}

func FuzzCollect(f *testing.F) {
	addRandomSeed(f, 16)

	f.Fuzz(func(
		t *testing.T,
		x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15 byte,
	) {
		xs := []byte{x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15}

		var x []Interval[elems.Uint8]

		for i, j := 0, len(xs); i < j; i += 2 {
			x = append(x, Range(elems.Uint8(xs[i]), elems.Uint8(xs[i+1])))
		}

		if w, z := plainUnion(x), Collect(x...); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("collect(x...) = %v", w)
			t.Logf("collect(x...) = %v (actual)", z)
			t.Fail()
		}
	})
}

func FuzzCollectInto(f *testing.F) {
	addRandomSeed(f, 16)

	f.Fuzz(func(
		t *testing.T,
		x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15 byte,
	) {
		xs := []byte{x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15}

		var x []Interval[elems.Uint8]

		for i, j := 0, len(xs); i < j; i += 2 {
			x = append(x, Range(elems.Uint8(xs[i]), elems.Uint8(xs[i+1])))
		}

		y := slices.Clone(x)

		if w, z := plainUnion(x), CollectInto(y, y...); !z.Equal(w) {
			t.Logf("x = %v", x)
			t.Logf("collect(x...) = %v", w)
			t.Logf("collect(x...) into(y) = %v (actual)", z)
			t.Fail()
		}
	})
}

func plainIsSubsetOf(x, y Set[elems.Uint8]) bool {
	var universe [256]bool

	for j, s := range []Set[elems.Uint8]{x, y} {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i] = j == 0
			}
		}
	}

	return !slices.Contains(universe[:], true)
}

func plainIntersection(sets ...Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]int

	for _, s := range sets {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i]++
			}
		}
	}

	n := len(sets)
	if n == 0 {
		return nil
	}

	return toSet(&universe, func(v int) bool { return v == n })
}

func plainUnion(sets ...Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]bool

	for _, s := range sets {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i] = true
			}
		}
	}

	return toSet(&universe, func(v bool) bool { return v })
}

func plainDifference(sets ...Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]bool

	for j, s := range sets {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i] = j == 0
			}
		}
	}

	return toSet(&universe, func(v bool) bool { return v })
}

func plainSymmetricDifference(sets ...Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]bool

	for _, s := range sets {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i] = !universe[i]
			}
		}
	}

	return toSet(&universe, func(v bool) bool { return v })
}

func toSet[T any](universe *[256]T, predicate func(T) bool) Set[elems.Uint8] {
	var x Set[elems.Uint8]

	lo := -1

	for i, v := range universe {
		if predicate(v) {
			if lo < 0 {
				lo = i
			}
		} else {
			if lo >= 0 {
				x = append(x, Range(elems.Uint8(lo), elems.Uint8(i)))
				lo = -1
			}
		}
	}

	return x
}
