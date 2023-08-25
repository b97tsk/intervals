package elems_test

import (
	"math"
	"net/netip"
	"testing"
	"time"

	"github.com/b97tsk/intervals/elems"
)

func TestPackage(t *testing.T) {
	t.Run("elems.Float32", func(t *testing.T) { testFloat[elems.Float32](t, math.SmallestNonzeroFloat32) })
	t.Run("elems.Float64", func(t *testing.T) { testFloat[elems.Float64](t, math.SmallestNonzeroFloat64) })
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
	t.Run("elems.Duration", testInteger[elems.Duration])
	t.Run("elems.Time", testTime)
	t.Run("elems.IP(v4)", func(t *testing.T) { testIP(t, "127.0.0.1", "127.0.0.2", "127.0.0.3") })
	t.Run("elems.IP(v6)", func(t *testing.T) { testIP(t, "::1", "::2", "::3") })
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Elem[E, U any] interface {
	Compare(E) int
	Next() E
	Unwrap() U
}

type FloatElem[E, U any] interface {
	Float
	Elem[E, U]
}

type IntegerElem[E, U any] interface {
	Integer
	Elem[E, U]
}

func testFloat[E FloatElem[E, U], U Float](t *testing.T, snz U) {
	var x E

	assert(t, x.Compare(x) == 0, "Compare didn't return 0.")
	assert(t, x.Compare(x.Next()) == -1, "Compare didn't return -1.")
	assert(t, x.Next().Compare(x) == +1, "Compare didn't return +1.")
	assert(t, x.Next().Unwrap() == snz, "Next didn't work.")
	assert(t, x.Next().Next().Unwrap() == snz*2, "Next twice didn't work.")
}

func testInteger[E IntegerElem[E, U], U Integer](t *testing.T) {
	var x E

	assert(t, x.Compare(x) == 0, "Compare didn't return 0.")
	assert(t, x.Compare(x.Next()) == -1, "Compare didn't return -1.")
	assert(t, x.Next().Compare(x) == +1, "Compare didn't return +1.")
	assert(t, x.Next().Unwrap() == 1, "Next didn't work.")
	assert(t, x.Next().Next().Unwrap() == 2, "Next twice didn't work.")
}

func testTime(t *testing.T) {
	x := elems.Time(time.Unix(0, 0))

	assert(t, x.Compare(x) == 0, "Compare didn't return 0.")
	assert(t, x.Compare(x.Next()) == -1, "Compare didn't return -1.")
	assert(t, x.Next().Compare(x) == +1, "Compare didn't return +1.")
	assert(t, x.Next().Unwrap().Compare(time.Unix(0, 1)) == 0, "Next didn't work.")
	assert(t, x.Next().Next().Unwrap().Compare(time.Unix(0, 2)) == 0, "Next twice didn't work.")
}

func testIP(t *testing.T, ip0, ip1, ip2 string) {
	x := elems.IP(netip.MustParseAddr(ip0))

	assert(t, x.Compare(x) == 0, "Compare didn't return 0.")
	assert(t, x.Compare(x.Next()) == -1, "Compare didn't return -1.")
	assert(t, x.Next().Compare(x) == +1, "Compare didn't return +1.")
	assert(t, x.Next().Unwrap().Compare(netip.MustParseAddr(ip1)) == 0, "Next didn't work.")
	assert(t, x.Next().Next().Unwrap().Compare(netip.MustParseAddr(ip2)) == 0, "Next twice didn't work.")
}

func assert(t *testing.T, ok bool, message string) {
	t.Helper()

	if ok {
		return
	}

	t.Fatal(message)
}
