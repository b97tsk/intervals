package elems

import "net/netip"

type IP netip.Addr

func (x IP) Compare(y IP) int { return x.Unwrap().Compare(y.Unwrap()) }

func (x IP) Next() IP { return IP(x.Unwrap().Next()) }

func (x IP) Unwrap() netip.Addr { return netip.Addr(x) }
