package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	// try run
	// gometalinter -j 1 --disable-all --enable=deadcode

	fmt.Println("Hello, Tokopedia!")
}
