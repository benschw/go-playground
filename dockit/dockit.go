package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var _ = log.Print

func main() {
	os.MkdirAll(DOCKIT_RUN_PATH, 0700)

	config := flag.String("config", "config.json", "json config")
	address := flag.String("address", "unix:///var/run/docker.sock", "docker address")
	service := flag.String("service", "", "service name")
	start := flag.Bool("start", false, "start service")
	stop := flag.Bool("stop", false, "stop service")
	flag.Parse()

	cfg, err := parseConfig(*config)
	if err != nil {
		panic(err)
	}

	lib := NewLib(cfg, *address)

	svcName := *service

	switch {
	case *start:
		if err = lib.Start(svcName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(svcName + " Started")
	case *stop:
		if err = lib.Stop(svcName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(svcName + " Stopped")
	}
}
