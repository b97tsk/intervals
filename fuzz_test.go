package intervals_test

import (
	"crypto/rand"
	"encoding/binary"
	"slices"
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func FuzzAddRange(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		z := append(Set[elems.Uint8](nil), x...)

		for _, r := range y {
			z.AddRange(r.Unwrap())
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

func FuzzDeleteRange(f *testing.F) {
	fuzz(f, func(t *testing.T, x, y Set[elems.Uint8]) {
		z := append(Set[elems.Uint8](nil), x...)

		for _, r := range y {
			z.DeleteRange(r.Unwrap())
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

func fuzz(f *testing.F, ff func(t *testing.T, x, y Set[elems.Uint8])) {
	data := make([]byte, 8*4)

	for i := 0; i < 10; i++ {
		if _, err := rand.Read(data); err != nil {
			f.Fatal(err)
		}

		f.Add(
			binary.LittleEndian.Uint64(data),
			binary.LittleEndian.Uint64(data[8:]),
			binary.LittleEndian.Uint64(data[16:]),
			binary.LittleEndian.Uint64(data[24:]),
		)
	}

	f.Fuzz(func(t *testing.T, x1, x2, y1, y2 uint64) {
		xs, ys := make([]byte, 8*2), make([]byte, 8*2)

		for i, u64 := range []uint64{x1, x2} {
			binary.LittleEndian.PutUint64(xs[i*8:], u64)
		}

		for i, u64 := range []uint64{y1, y2} {
			binary.LittleEndian.PutUint64(ys[i*8:], u64)
		}

		slices.Sort(xs)
		slices.Sort(ys)

		var x, y Set[elems.Uint8]

		for i, j := 0, len(xs); i < j; i += 2 {
			if lo, hi := elems.Uint8(xs[i]), elems.Uint8(xs[i+1]); lo < hi {
				x.AddRange(lo, hi)
			}
		}

		for i, j := 0, len(ys); i < j; i += 2 {
			if lo, hi := elems.Uint8(ys[i]), elems.Uint8(ys[i+1]); lo < hi {
				y.AddRange(lo, hi)
			}
		}

		ff(t, x, y)
	})
}

func plainIsSubsetOf(x, y Set[elems.Uint8]) bool {
	var universe [256]uint8

	for x, s := range []Set[elems.Uint8]{x, y} {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i] = uint8(x + 1)
			}
		}
	}

	return !slices.Contains(universe[:], 1)
}

func plainIntersection(x, y Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]uint8

	for _, s := range []Set[elems.Uint8]{x, y} {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i]++
			}
		}
	}

	var z Set[elems.Uint8]

	lo := -1

	for i, v := range universe {
		if v == 2 {
			if lo < 0 {
				lo = i
			}
		} else {
			if lo >= 0 {
				z = append(z, Interval[elems.Uint8]{elems.Uint8(lo), elems.Uint8(i)})
				lo = -1
			}
		}
	}

	return z
}

func plainUnion(x, y Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]uint8

	for _, s := range []Set[elems.Uint8]{x, y} {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i]++
			}
		}
	}

	var z Set[elems.Uint8]

	lo := -1

	for i, v := range universe {
		if v > 0 {
			if lo < 0 {
				lo = i
			}
		} else {
			if lo >= 0 {
				z = append(z, Interval[elems.Uint8]{elems.Uint8(lo), elems.Uint8(i)})
				lo = -1
			}
		}
	}

	return z
}

func plainDifference(x, y Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]uint8

	for x, s := range []Set[elems.Uint8]{x, y} {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i] = uint8(x + 1)
			}
		}
	}

	var z Set[elems.Uint8]

	lo := -1

	for i, v := range universe {
		if v == 1 {
			if lo < 0 {
				lo = i
			}
		} else {
			if lo >= 0 {
				z = append(z, Interval[elems.Uint8]{elems.Uint8(lo), elems.Uint8(i)})
				lo = -1
			}
		}
	}

	return z
}

func plainSymmetricDifference(x, y Set[elems.Uint8]) Set[elems.Uint8] {
	var universe [256]uint8

	for _, s := range []Set[elems.Uint8]{x, y} {
		for _, r := range s {
			for i := r.Low; i < r.High; i++ {
				universe[i]++
			}
		}
	}

	var z Set[elems.Uint8]

	lo := -1

	for i, v := range universe {
		if v == 1 {
			if lo < 0 {
				lo = i
			}
		} else {
			if lo >= 0 {
				z = append(z, Interval[elems.Uint8]{elems.Uint8(lo), elems.Uint8(i)})
				lo = -1
			}
		}
	}

	return z
}
