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
}
