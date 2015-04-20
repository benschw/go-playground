package main

import (
	"fmt"
	"testing"

	"github.com/benschw/dns-clb-go/dns"
	"github.com/benschw/opin-go/rando"
)

type StaticAddressGetter struct {
	Val dns.Address
}

func (lb *StaticAddressGetter) GetAddress(address string) (dns.Address, error) {
	return lb.Val, nil
}

func TestDemoEndpoint(t *testing.T) {
	// given

	expectedGreeting := "\"hello world\""

	host := "localhost"
	port := uint16(rando.Port())

	go RunServer(fmt.Sprintf("%s:%d", host, port))

	client := GreetingClient{
		Lb: &StaticAddressGetter{Val: dns.Address{Address: host, Port: port}},
	}

	// when
	greeting, _ := client.GetGreeting()

	// then
	if expectedGreeting != string(greeting[:]) {
		t.Errorf("expected '%s', got '%s'", expectedGreeting, greeting)
	}

}