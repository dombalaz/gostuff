package main

import "fmt"

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Doubler interface {
	Integer | Float
}

func Double[T Doubler](n T) T {
	return n * 2
}

func main() {
	var n1 int = 1
	var n2 float32 = 2.5
	var n3 float64 = 2.5
	var n4 int8 = 10
	var n5 int64 = 5

	fmt.Println(Double(n1))
	fmt.Println(Double(n2))
	fmt.Println(Double(n3))
	fmt.Println(Double(n4))
	fmt.Println(Double(n5))
}
