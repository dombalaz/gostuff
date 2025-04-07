package main

import "fmt"

func launch() {
	f := func(nums chan<- int, x int) {
		for i := range 10 {
			nums <- i * x
		}
		close(nums)
	}

	c1 := make(chan int)
	c2 := make(chan int)

	go f(c1, 1)
	go f(c2, 100)

	for count := 2; count != 0; {
		select {
		case n, ok := <-c1:
			if !ok {
				c1 = nil
				count--
				break
			}
			fmt.Printf("g1: %d\n", n)
		case n, ok := <-c2:
			if !ok {
				c2 = nil
				count--
				break
			}
			fmt.Printf("g2: %d\n", n)
		}
	}
}

func main() {
	launch()
}
