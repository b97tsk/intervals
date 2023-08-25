package elems

import (
	"cmp"
	"time"
)

type Duration time.Duration

func (x Duration) Compare(y Duration) int { return cmp.Compare(x, y) }

func (x Duration) Next() Duration { return Duration(x.Unwrap() + time.Nanosecond) }

func (x Duration) Unwrap() time.Duration { return time.Duration(x) }

type Time time.Time

func (x Time) Compare(y Time) int { return x.Unwrap().Compare(y.Unwrap()) }

func (x Time) Next() Time { return Time(x.Unwrap().Add(time.Nanosecond)) }

func (x Time) Unwrap() time.Time { return time.Time(x) }
