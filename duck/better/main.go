package main

import (
	"fmt"
	"log"

	"github.com/benschw/dns-clb-go/clb"
	"github.com/benschw/dns-clb-go/dns"
)

type AddressGetter interface {
	GetAddress(string) (dns.Address, error)
}

func Discover(lb AddressGetter, addr string) string {
	address, _ := lb.GetAddress(addr)

	return fmt.Sprintf("Discovered Address: '%s'", address)
}

func main() {
	lb := clb.NewClb("8.8.8.8", "53", clb.Random)

	log.Println(Discover(lb, "foo.service.fliglio.com"))
}
