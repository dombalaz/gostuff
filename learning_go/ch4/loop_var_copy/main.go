package main

import "fmt"

// each iteration the i and v are new variables, i.e. have different memory address on each iteration
// this applies to go version 1.22+
func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Printf("i: %p, v: %p\n", &i, &v)
	}
}
