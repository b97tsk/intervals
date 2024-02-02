package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestIntersection(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Set[E]{{5, 7}}.Intersection(Set[E]{{1, 3}}),
			Set[E]{},
		},
		{
			Set[E]{{3, 7}}.Intersection(Set[E]{{1, 5}}),
			Set[E]{{3, 5}},
		},
		{
			Set[E]{{3, 11}, {13, 21}}.Intersection(Set[E]{{1, 5}, {9, 15}, {19, 23}}),
			Set[E]{{3, 5}, {9, 11}, {13, 15}, {19, 21}},
		},
		{Reduce(Intersection[E]), Set[E]{}},
		{Reduce(Intersection, Set[E]{}), Set[E]{}},
		{
			func() Set[E] {
				var x2, x3, x5 Set[E]

				for i := 2; i < 100; i += 2 {
					x2 = Add(x2, Unit(E(i)))
				}

				for i := 3; i < 100; i += 3 {
					x3 = Add(x3, Unit(E(i)))
				}

				for i := 5; i < 100; i += 5 {
					x5 = Add(x5, Unit(E(i)))
				}

				return Reduce(Intersection, x2, x3, x5)
			}(),
			Set[E]{{30, 31}, {60, 61}, {90, 91}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}
