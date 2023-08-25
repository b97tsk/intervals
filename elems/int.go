package elems

import "cmp"

type Int int

func (x Int) Compare(y Int) int { return cmp.Compare(x, y) }

func (x Int) Next() Int { return x + 1 }

func (x Int) Unwrap() int { return int(x) }

type Int8 int8

func (x Int8) Compare(y Int8) int { return cmp.Compare(x, y) }

func (x Int8) Next() Int8 { return x + 1 }

func (x Int8) Unwrap() int8 { return int8(x) }

type Int16 int16

func (x Int16) Compare(y Int16) int { return cmp.Compare(x, y) }

func (x Int16) Next() Int16 { return x + 1 }

func (x Int16) Unwrap() int16 { return int16(x) }

type Int32 int32

func (x Int32) Compare(y Int32) int { return cmp.Compare(x, y) }

func (x Int32) Next() Int32 { return x + 1 }

func (x Int32) Unwrap() int32 { return int32(x) }

type Int64 int64

func (x Int64) Compare(y Int64) int { return cmp.Compare(x, y) }

func (x Int64) Next() Int64 { return x + 1 }

func (x Int64) Unwrap() int64 { return int64(x) }
