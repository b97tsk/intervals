package elems

import "cmp"

type Float32 float32

func (x Float32) Compare(y Float32) int { return cmp.Compare(x, y) }

func (x Float32) Unwrap() float32 { return float32(x) }

type Float64 float64

func (x Float64) Compare(y Float64) int { return cmp.Compare(x, y) }

func (x Float64) Unwrap() float64 { return float64(x) }
