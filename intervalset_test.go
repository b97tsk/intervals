package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestCreation(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Range[E](1, 5).Set(),
			Set[E]{{1, 5}},
		},
		{
			Range[E](5, 1).Set(),
			nil,
		},
		{
			Collect(Range[E](1, 5), Range[E](7, 11), Range[E](13, 17)),
			Set[E]{{1, 5}, {7, 11}, {13, 17}},
		},
		{
			Collect(Range[E](13, 17), Range[E](7, 11), Range[E](1, 5)),
			Set[E]{{1, 5}, {7, 11}, {13, 17}},
		},
		{
			Collect(Range[E](1, 7), Range[E](5, 13), Range[E](11, 17)),
			Set[E]{{1, 17}},
		},
		{
			Collect[E](),
			nil,
		},
		{
			Collect(Range[E](5, 1), Range[E](3, 7), Range[E](11, 9), Range[E](5, 1)),
			Set[E]{{3, 7}},
		},
	}

	for i, c := range testCases {
		if !c.Actual.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Actual)
		}
	}
}

func TestAdd(t *testing.T) {
	type E = elems.Int

	add := func(s Set[E], r Interval[E]) Set[E] {
		s.Add(r)
		return s
	}

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			add(Set[E]{{1, 4}, {9, 12}}, Range[E](5, 8)),
			Set[E]{{1, 4}, {5, 8}, {9, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, One[E](6)),
			Set[E]{{1, 4}, {6, 7}, {9, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, Range[E](4, 8)),
			Set[E]{{1, 8}, {9, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, Range[E](5, 9)),
			Set[E]{{1, 4}, {5, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, Range[E](4, 9)),
			Set[E]{{1, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, One[E](10)),
			Set[E]{{1, 4}, {9, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, Range[E](9, 12)),
			Set[E]{{1, 4}, {9, 12}},
		},
		{
			add(Set[E]{{1, 4}, {9, 12}}, Range[E](12, 9)),
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

	del := func(s Set[E], r Interval[E]) Set[E] {
		s.Delete(r)
		return s
	}

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](7, 10)),
			Set[E]{{1, 4}, {13, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](7, 9)),
			Set[E]{{1, 4}, {9, 10}, {13, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](8, 10)),
			Set[E]{{1, 4}, {7, 8}, {13, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, One[E](8)),
			Set[E]{{1, 4}, {7, 8}, {9, 10}, {13, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](1, 16)),
			Set[E]{},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](1, 15)),
			Set[E]{{15, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](2, 16)),
			Set[E]{{1, 2}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, One[E](5)),
			Set[E]{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](4, 7)),
			Set[E]{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			del(Set[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E](7, 4)),
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
		s.ContainsOne(0) == false,
		s.ContainsOne(1) == true,
		s.ContainsOne(2) == true,
		s.ContainsOne(3) == false,
		s.ContainsOne(4) == false,
		s.ContainsOne(5) == true,
		s.ContainsOne(6) == true,
		s.ContainsOne(7) == false,
		s.Contains(Range[E](1, 3)) == true,
		s.Contains(Range[E](3, 5)) == false,
		s.Contains(Range[E](5, 7)) == true,
		s.Contains(Range[E](1, 7)) == false,
		s.Contains(Range[E](1, 1)) == false,
		s.Contains(Range[E](2, 2)) == false,
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
