package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestUnion(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Set[E]{{1, 3}, {5, 7}}.Union(Set[E]{}),
			Set[E]{{1, 3}, {5, 7}},
		},
		{
			Set[E]{}.Union(Set[E]{{1, 3}, {5, 7}}),
			Set[E]{{1, 3}, {5, 7}},
		},
		{
			Set[E]{{5, 7}}.Union(Set[E]{{1, 3}}),
			Set[E]{{1, 3}, {5, 7}},
		},
		{
			Set[E]{{3, 7}}.Union(Set[E]{{1, 5}}),
			Set[E]{{1, 7}},
		},
		{
			Set[E]{{3, 11}, {13, 21}}.Union(Set[E]{{1, 5}, {9, 15}, {19, 23}}),
			Set[E]{{1, 23}},
		},
		{Union[E](), Set[E]{}},
		{Union(Set[E]{}), Set[E]{}},
		{
			func() Set[E] {
				var x2, x3, x5 Set[E]

				for i := 2; i < 32; i += 2 {
					x2.Add(E(i))
				}

				for i := 3; i < 32; i += 3 {
					x3.Add(E(i))
				}

				for i := 5; i < 32; i += 5 {
					x5.Add(E(i))
				}

				return Union(x2, x3, x5)
			}(),
			Set[E]{{2, 7}, {8, 11}, {12, 13}, {14, 17}, {18, 19}, {20, 23}, {24, 29}, {30, 31}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}
