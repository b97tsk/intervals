package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestOverlaps(t *testing.T) {
	type E = elems.Int

	assertions := []bool{
		Set[E]{}.Overlaps(Set[E]{}) == false,
		Set[E]{{5, 7}}.Overlaps(Set[E]{{1, 3}}) == false,
		Set[E]{{3, 7}}.Overlaps(Set[E]{{1, 5}}) == true,
		Set[E]{{5, 7}}.Overlaps(Set[E]{{1, 3}, {9, 11}}) == false,
		Set[E]{{3, 9}}.Overlaps(Set[E]{{1, 5}, {7, 11}}) == true,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
