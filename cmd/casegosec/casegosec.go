package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// try run
	// gometalinter -j 1 --disable-all --enable=gosec

	foo := "foo"
	fileName := fmt.Sprintf("%s bar", foo)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(file))
}
