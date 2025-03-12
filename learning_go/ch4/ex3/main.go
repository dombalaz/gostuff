package main

import "fmt"

func main() {
	var total int
	for i := range 10 {
		total := total + i // shadowed variable
		fmt.Println(total) // prints 0+i
	}
	fmt.Println(total) // prints 0

	total = 0
	for i := range 10 {
		total += i         // accumulate
		fmt.Println(total) // prints total+i
	}
	fmt.Println(total)
}
