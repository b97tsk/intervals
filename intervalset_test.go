package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestAdd(t *testing.T) {
	type E = elems.Int

	addSingle := func(s Set[E], v E) Set[E] {
		s.Add(v)
		return s
	}
	addInterval := func(s Set[E], r Interval[E]) Set[E] {
		s.AddRange(r.Unwrap())
		return s
	}

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			addInterval(Set[E]{{1, 4}, {9, 12}}, Interval[E]{5, 8}),
			Set[E]{{1, 4}, {5, 8}, {9, 12}},
		},
		{
			addSingle(Set[E]{{1, 4}, {9, 12}}, 6),
			Set[E]{{1, 4}, {6, 7}, {9, 12}},
		},
		{
			addInterval(Set[E]{{1, 4}, {9, 12}}, Interval[E]{4, 8}),
			Set[E]{{1, 8}, {9, 12}},
		},
		{
			addInterval(Set[E]{{1, 4}, {9, 12}}, Interval[E]{5, 9}),
			Set[E]{{1, 4}, {5, 12}},
		},
		{
			addInterval(Set[E]{{1, 4}, {9, 12}}, Interval[E]{4, 9}),
			Set[E]{{1, 12}},
		},
		{
			addSingle(Set[E]{{1, 4}, {9, 12}}, 10),
			Set[E]{{1, 4}, {9, 12}},
		},
		{
			addInterval(Set[E]{{1, 4}, {9, 12}}, Interval[E]{9, 12}),
			Set[E]{{1, 4}, {9, 12}},
		},
		{
			addInterval(Set[E]{{1, 4}, {9, 12}}, Interval[E]{12, 9}),
			Set[E]{{1, 4}, {9, 12}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}

func TestDelete(t *testing.T) {
	type E = elems.Int

	deleteSingle := func(s Set[E], v E) Set[E] {
		s.Delete(v)
		return s
	}
	deleteInterval := func(s Set[E], r Interval[E]) Set[E] {
		s.DeleteRange(r.Unwrap())
		return s
	}

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{7, 10}),
			Set[E]{{1, 4}, {13, 16}},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{7, 9}),
			Set[E]{{1, 4}, {9, 10}, {13, 16}},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{8, 10}),
			Set[E]{{1, 4}, {7, 8}, {13, 16}},
		},
		{
			deleteSingle(Set[E]{{1, 4}, {7, 10}, {13, 16}}, 8),
			Set[E]{{1, 4}, {7, 8}, {9, 10}, {13, 16}},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{1, 16}),
			Set[E]{},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{1, 15}),
			Set[E]{{15, 16}},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{2, 16}),
			Set[E]{{1, 2}},
		},
		{
			deleteSingle(Set[E]{{1, 4}, {7, 10}, {13, 16}}, 5),
			Set[E]{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{4, 7}),
			Set[E]{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			deleteInterval(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Interval[E]{7, 4}),
			Set[E]{{1, 4}, {7, 10}, {13, 16}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}

func TestContains(t *testing.T) {
	type E = elems.Int

	s := Set[E]{{1, 3}, {5, 7}}

	assertions := []bool{
		s.Contains(0) == false,
		s.Contains(1) == true,
		s.Contains(2) == true,
		s.Contains(3) == false,
		s.Contains(4) == false,
		s.Contains(5) == true,
		s.Contains(6) == true,
		s.Contains(7) == false,
		s.ContainsRange(1, 3) == true,
		s.ContainsRange(3, 5) == false,
		s.ContainsRange(5, 7) == true,
		s.ContainsRange(1, 7) == false,
		s.ContainsRange(1, 1) == false,
		s.ContainsRange(2, 2) == false,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestEqual(t *testing.T) {
	type E = elems.Int

	assertions := []bool{
		Set[E]{{1, 3}, {5, 7}}.Equal(Set[E]{{1, 3}, {5, 7}}),
		!Set[E]{{1, 3}, {5, 7}}.Equal(Set[E]{{1, 3}, {5, 9}}),
		!Set[E]{{1, 3}, {5, 7}}.Equal(Set[E]{{1, 3}}),
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestExtent(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Interval[E]
	}{
		{
			Set[E]{{1, 3}, {5, 7}}.Extent(),
			Interval[E]{1, 7},
		},
		{
			Set[E]{}.Extent(),
			Interval[E]{},
		},
	}
	for i, c := range testCases {
		if c.Actual != c.Expected {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}

func TestIsSubsetOf(t *testing.T) {
	type E = elems.Int

	assertions := []bool{
		Set[E]{}.IsSubsetOf(Set[E]{}) == true,
		Set[E]{{3, 9}}.IsSubsetOf(Set[E]{{1, 11}}) == true,
		Set[E]{{3, 9}}.IsSubsetOf(Set[E]{{1, 5}, {7, 11}}) == false,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
