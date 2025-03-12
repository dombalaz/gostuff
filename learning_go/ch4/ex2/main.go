package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := make([]int, 0, 100)
	for range 100 {
		r = append(r, rand.Intn(100))
	}
	fmt.Println(r)

	for _, v := range r {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%3 == 0:
			fmt.Println("Three!")
		case v%2 == 0:
			fmt.Println("Two!")
		default:
			fmt.Println("Never mind")
		}
	}
}
