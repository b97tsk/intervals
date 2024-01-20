package intervals_test

import (
	"testing"

	. "github.com/b97tsk/intervals"
	"github.com/b97tsk/intervals/elems"
)

func TestIsSubsetOf(t *testing.T) {
	type E = elems.Int

	assertions := []bool{
		Set[E]{}.IsSubsetOf(Set[E]{}) == true,
		Set[E]{{3, 7}}.IsSubsetOf(Set[E]{{1, 9}}) == true,
		Set[E]{{3, 7}}.IsSubsetOf(Set[E]{{1, 5}}) == false,
		Set[E]{{3, 7}}.IsSubsetOf(Set[E]{{9, 11}}) == false,
		Set[E]{{7, 9}}.IsSubsetOf(Set[E]{{1, 3}, {5, 11}}) == true,
		Set[E]{{3, 9}}.IsSubsetOf(Set[E]{{1, 5}, {7, 11}}) == false,
		Set[E]{{3, 5}, {7, 11}}.IsSubsetOf(Set[E]{{1, 9}}) == false,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
