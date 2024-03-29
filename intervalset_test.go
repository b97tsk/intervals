package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestIsValid(t *testing.T) {
	type E = elems.Int

	assertions := []bool{
		Interval[E]{}.IsValid(),
		Unit[E](0).IsValid(),
		Range[E](1, 5).IsValid(),
		!Range[E](5, 1).IsValid(),
		!Range[E](5, 5).IsValid(),
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestCreation(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Unit[E](1).Set(),
			Set[E]{{1, 2}},
		},
		{
			Range[E](1, 5).Set(),
			Set[E]{{1, 5}},
		},
		{
			Interval[E]{}.Set(),
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

	shouldPanic(t, func() { _ = Range[E](5, 1).Set() }, "Range[E](5, 1).Set() didn't panic.")
	shouldPanic(t, func() { _ = Range[E](5, 5).Set() }, "Range[E](5, 5).Set() didn't panic.")
}

func TestAdd(t *testing.T) {
	type E = elems.Int

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Add(Set[E]{{1, 5}, {11, 15}}, Range[E](5, 11)),
			Set[E]{{1, 15}},
		},
		{
			Add(Set[E]{{1, 5}, {11, 15}}, Range[E](5, 9)),
			Set[E]{{1, 9}, {11, 15}},
		},
		{
			Add(Set[E]{{1, 5}, {11, 15}}, Range[E](7, 11)),
			Set[E]{{1, 5}, {7, 15}},
		},
		{
			Add(Set[E]{{1, 5}, {11, 15}}, Range[E](7, 9)),
			Set[E]{{1, 5}, {7, 9}, {11, 15}},
		},
		{
			Add(Set[E]{{1, 5}, {11, 15}}, Range[E](9, 7)),
			Set[E]{{1, 5}, {11, 15}},
		},
		{
			Add(Set[E]{{1, 5}, {11, 15}}, Range[E](15, 11)),
			Set[E]{{1, 5}, {11, 15}},
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

	testCases := []struct {
		Actual, Expected Set[E]
	}{
		{
			Delete(Set[E]{{1, 3}, {5, 11}, {13, 15}}, Range[E](5, 11)),
			Set[E]{{1, 3}, {13, 15}},
		},
		{
			Delete(Set[E]{{1, 3}, {5, 11}, {13, 15}}, Range[E](5, 9)),
			Set[E]{{1, 3}, {9, 11}, {13, 15}},
		},
		{
			Delete(Set[E]{{1, 3}, {5, 11}, {13, 15}}, Range[E](7, 11)),
			Set[E]{{1, 3}, {5, 7}, {13, 15}},
		},
		{
			Delete(Set[E]{{1, 3}, {5, 11}, {13, 15}}, Range[E](7, 9)),
			Set[E]{{1, 3}, {5, 7}, {9, 11}, {13, 15}},
		},
		{
			Delete(Set[E]{{1, 3}, {5, 11}, {13, 15}}, Range[E](9, 7)),
			Set[E]{{1, 3}, {5, 11}, {13, 15}},
		},
		{
			Delete(Set[E]{{1, 3}, {5, 11}, {13, 15}}, Range[E](5, 3)),
			Set[E]{{1, 3}, {5, 11}, {13, 15}},
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
		s.ContainsUnit(0) == false,
		s.ContainsUnit(1) == true,
		s.ContainsUnit(2) == true,
		s.ContainsUnit(3) == false,
		s.ContainsUnit(4) == false,
		s.ContainsUnit(5) == true,
		s.ContainsUnit(6) == true,
		s.ContainsUnit(7) == false,
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
			Range[E](1, 7),
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

func shouldPanic(t *testing.T, f func(), name string) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Log(name, "did not panic.")
			t.Fail()
		}
	}()

	f()
}
