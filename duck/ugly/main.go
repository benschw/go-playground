package main

import (
	"fmt"
	"log"

	"github.com/benschw/dns-clb-go/clb"
)

func Discover(lb clb.LoadBalancer, addr string) string {
	address, _ := lb.GetAddress(addr)

	return fmt.Sprintf("Discovered Address: '%s'", address)
}

func main() {
	lb := clb.NewClb("8.8.8.8", "53", clb.Random)

	log.Println(Discover(lb, "foo.service.fliglio.com"))
}
