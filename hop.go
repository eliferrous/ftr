package ftr

import (
	"net/netip"
)

type Hop struct {
	Count      int
	Host       string
	Addr       *netip.Addr
	Unanswered bool
	ASN        string

	// rtt stats in ms
	Ploss float64
	Sent  int
	Last  float64
	Avg   float64
	Best  float64
	Worst float64
	Stdev float64
}
