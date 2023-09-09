package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestDifference(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Set[E]{{5, 7}}.Difference(Set[E]{{1, 3}}),
			Set[E]{{5, 7}},
		},
		{
			Set[E]{{3, 7}}.Difference(Set[E]{{1, 5}}),
			Set[E]{{5, 7}},
		},
		{
			Set[E]{{3, 11}, {13, 25}}.Difference(Set[E]{{1, 5}, {9, 15}, {19, 23}}),
			Set[E]{{5, 9}, {15, 19}, {23, 25}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{9, 21}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{9, 19}}),
			Set[E]{{1, 5}, {19, 21}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{9, 23}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{9, 25}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{5, 21}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{7, 21}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{11, 21}}),
			Set[E]{{1, 5}, {9, 11}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{7, 23}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{
			Set[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}}.Difference(Set[E]{{5, 25}}),
			Set[E]{{1, 5}, {25, 29}},
		},
		{Difference[E](), Set[E]{}},
		{Difference(Set[E]{}), Set[E]{}},
		{
			func() Set[E] {
				var x2, x3, x5 Set[E]

				for i := 2; i < 32; i += 2 {
					x2.Add(One(E(i)))
				}

				for i := 3; i < 32; i += 3 {
					x3.Add(One(E(i)))
				}

				for i := 5; i < 32; i += 5 {
					x5.Add(One(E(i)))
				}

				return Difference(x2, x3, x5)
			}(),
			Set[E]{{2, 3}, {4, 5}, {8, 9}, {14, 15}, {16, 17}, {22, 23}, {26, 27}, {28, 29}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}
