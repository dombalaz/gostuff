package main

import (
	"fmt"
	"math"
	"sync"
)

func sqrtMap() map[int]float64 {
	r := make(map[int]float64)
	for i := range 100_000 {
		r[i] = math.Sqrt(float64(i))
	}
	return r
}

func main() {
	f := sync.OnceValue(sqrtMap)

	for i := 0; i < 100_000; i += 1_000 {
		fmt.Printf("%d: %f\n", i, f()[i])
	}
}
