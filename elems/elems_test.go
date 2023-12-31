package elems_test

import (
	"testing"

	"github.com/b97tsk/intervals/elems"
)

func TestPackage(t *testing.T) {
	t.Run("elems.Float32", testFloat[elems.Float32])
	t.Run("elems.Float64", testFloat[elems.Float64])
	t.Run("elems.Int", testInteger[elems.Int])
	t.Run("elems.Int8", testInteger[elems.Int8])
	t.Run("elems.Int16", testInteger[elems.Int16])
	t.Run("elems.Int32", testInteger[elems.Int32])
	t.Run("elems.Int64", testInteger[elems.Int64])
	t.Run("elems.Uint", testInteger[elems.Uint])
	t.Run("elems.Uint8", testInteger[elems.Uint8])
	t.Run("elems.Uint16", testInteger[elems.Uint16])
	t.Run("elems.Uint32", testInteger[elems.Uint32])
	t.Run("elems.Uint64", testInteger[elems.Uint64])
	t.Run("elems.Uintptr", testInteger[elems.Uintptr])
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type FloatElem[E, U any] interface {
	Float
	Compare(E) int
	Unwrap() U
}

type IntegerElem[E, U any] interface {
	Integer
	Compare(E) int
	Next() E
	Unwrap() U
}

func testFloat[E FloatElem[E, U], U Float](t *testing.T) {
	var x E

	assert(t, x.Compare(x) == 0, "Compare didn't return 0.")
	assert(t, x.Compare(x+1) == -1, "Compare didn't return -1.")
	assert(t, (x+1).Compare(x) == +1, "Compare didn't return +1.")
	assert(t, x.Unwrap() == U(0), "Unwrap didn't work.")
}

func testInteger[E IntegerElem[E, U], U Integer](t *testing.T) {
	var x E

	assert(t, x.Compare(x) == 0, "Compare didn't return 0.")
	assert(t, x.Compare(x.Next()) == -1, "Compare didn't return -1.")
	assert(t, x.Next().Compare(x) == +1, "Compare didn't return +1.")
	assert(t, x.Next().Unwrap() == 1, "Next didn't work.")
	assert(t, x.Next().Next().Unwrap() == 2, "Next twice didn't work.")
}

func assert(t *testing.T, ok bool, message string) {
	t.Helper()

	if !ok {
		t.Fatal(message)
	}
}
