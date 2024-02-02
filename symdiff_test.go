package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestSymmetricDifference(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Set[E]{{5, 7}}.SymmetricDifference(Set[E]{{1, 3}}),
			Set[E]{{1, 3}, {5, 7}},
		},
		{
			Set[E]{{3, 7}}.SymmetricDifference(Set[E]{{1, 5}}),
			Set[E]{{1, 3}, {5, 7}},
		},
		{
			Set[E]{{3, 11}, {13, 25}}.SymmetricDifference(Set[E]{{1, 5}, {9, 15}, {19, 23}}),
			Set[E]{{1, 3}, {5, 9}, {11, 13}, {15, 19}, {23, 25}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{9, 21}}),
			Set[E]{{1, 5}, {13, 17}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{9, 19}}),
			Set[E]{{1, 5}, {13, 17}, {19, 21}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{9, 23}}),
			Set[E]{{1, 5}, {13, 17}, {21, 23}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{9, 25}}),
			Set[E]{{1, 5}, {13, 17}, {21, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{5, 21}}),
			Set[E]{{1, 9}, {13, 17}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{7, 21}}),
			Set[E]{{1, 5}, {7, 9}, {13, 17}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{11, 21}}),
			Set[E]{{1, 5}, {9, 11}, {13, 17}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{7, 23}}),
			Set[E]{{1, 5}, {7, 9}, {13, 17}, {21, 23}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.SymmetricDifference(Set[E]{{5, 25}}),
			Set[E]{{1, 9}, {13, 17}, {21, 29}},
		},
		{Reduce(SymmetricDifference[E]), Set[E]{}},
		{Reduce(SymmetricDifference, Set[E]{}), Set[E]{}},
		{
			func() Set[E] {
				var x2, x3, x5 Set[E]

				for i := 2; i < 32; i += 2 {
					x2 = Add(x2, Unit(E(i)))
				}

				for i := 3; i < 32; i += 3 {
					x3 = Add(x3, Unit(E(i)))
				}

				for i := 5; i < 32; i += 5 {
					x5 = Add(x5, Unit(E(i)))
				}

				return Reduce(SymmetricDifference, x2, x3, x5)
			}(),
			Set[E]{{2, 6}, {8, 10}, {14, 15}, {16, 17}, {21, 23}, {25, 29}, {30, 31}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}
