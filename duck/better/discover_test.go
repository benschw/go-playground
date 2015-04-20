package main

import (
	"testing"

	"github.com/benschw/dns-clb-go/dns"
)

type StaticAddressGetter struct {
	Val dns.Address
}

func (lb *StaticAddressGetter) GetAddress(address string) (dns.Address, error) {
	return lb.Val, nil
}

func TestDiscover(t *testing.T) {
	//given
	expected := "Discovered Address: 'foo:8080'"
	lb := &StaticAddressGetter{Val: dns.Address{Address: "foo", Port: 8080}}

	// when
	found := Discover(lb, "")

	// then
	if found != expected {
		t.Errorf("\"%s\" not equal to \"%s\"", found, expected)
	}
}
