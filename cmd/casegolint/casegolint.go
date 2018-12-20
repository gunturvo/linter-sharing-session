package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func main() {
	// try run
	// gometalinter -j 1 --disable-all --enable=golint

	fmt.Println(Add(42, 13))
}
