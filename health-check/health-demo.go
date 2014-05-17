package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "http service address") // Q=17, R=18

func status(w http.ResponseWriter, req *http.Request) {
	health := "OK"

	if _, err := os.Stat("/tmp/fail-healthcheck"); err == nil {
		//file exists
		health = "FAIL"
	}
	fmt.Fprintf(w, "%+v", health)
}

func main() {
	flag.Parse()
	http.Handle("/status", http.HandlerFunc(status))

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
