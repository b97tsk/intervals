package elems

import "cmp"

type Uint uint

func (x Uint) Compare(y Uint) int { return cmp.Compare(x, y) }

func (x Uint) Next() Uint { return x + 1 }

func (x Uint) Unwrap() uint { return uint(x) }

type Uint8 uint8

func (x Uint8) Compare(y Uint8) int { return cmp.Compare(x, y) }

func (x Uint8) Next() Uint8 { return x + 1 }

func (x Uint8) Unwrap() uint8 { return uint8(x) }

type Uint16 uint16

func (x Uint16) Compare(y Uint16) int { return cmp.Compare(x, y) }

func (x Uint16) Next() Uint16 { return x + 1 }

func (x Uint16) Unwrap() uint16 { return uint16(x) }

type Uint32 uint32

func (x Uint32) Compare(y Uint32) int { return cmp.Compare(x, y) }

func (x Uint32) Next() Uint32 { return x + 1 }

func (x Uint32) Unwrap() uint32 { return uint32(x) }

type Uint64 uint64

func (x Uint64) Compare(y Uint64) int { return cmp.Compare(x, y) }

func (x Uint64) Next() Uint64 { return x + 1 }

func (x Uint64) Unwrap() uint64 { return uint64(x) }

type Uintptr uintptr

func (x Uintptr) Compare(y Uintptr) int { return cmp.Compare(x, y) }

func (x Uintptr) Next() Uintptr { return x + 1 }

func (x Uintptr) Unwrap() uintptr { return uintptr(x) }
