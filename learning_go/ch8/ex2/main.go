package main

import (
	"fmt"
)

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyInt int

func (v MyInt) String() string {
	// usage of %v causes recursion as it calls String() again
	return fmt.Sprintf("%d", v)
}

type MyFloat float64

func (v MyFloat) String() string {
	return fmt.Sprintf("%f", v)
}

func Print[T Printable](v T) {
	fmt.Println(v)
}

func main() {
	Print(MyInt(10))
	Print(MyFloat(10.5))
}
