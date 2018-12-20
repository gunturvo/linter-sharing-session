package main

import "fmt"

func a() error {
	fmt.Println("this function returns an error") // EXCLUDED
	return nil
}

func main() {
	// try run
	// gometalinter -j 1 --disable-all --enable=errcheck

	// Single error return
	_ = a() // BLANK
	a()     // UNCHECKED
}
