package main

import (
	"log"
	"fmt"
	"launchpad.net/goyaml"
	"io/ioutil"
)


var data = `
foo: bar
baz:
  c: 2
  d: [3, 4]
`
type T struct {
        Foo string
        Baz struct{C int; D []int ",flow"}
}


func main() {
	t := T{}

	err := goyaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t)


	fileData, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	t2 := T{}

	err = goyaml.Unmarshal(fileData, &t2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t2)


}