package main

import "fmt"

func justPrint(firstCategory, secondCategory string) (string, string) {
	if firstCategory == "station" {
		if secondCategory == "station" {
			return firstCategory, secondCategory
		} else if secondCategory == "city" {
			return firstCategory, secondCategory

		}
	} else if firstCategory == "city" {
		if secondCategory == "station" {
			return firstCategory, secondCategory

		} else if secondCategory == "city" {
			return firstCategory, secondCategory

		}
	}
	return "", ""
}

func main() {
	// try run
	// gometalinter -j 1 --disable-all --enable=goconst

	catOne := "lorem"
	catTwo := "ipsum"

	fmt.Println(justPrint(catOne, catTwo))
}
