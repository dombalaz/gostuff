package main

import (
	"fmt"
	"time"
)

func defaultChan(ch1 chan int, ch2 chan int) {
	// Okay to do both
	val := <-ch1
	fmt.Println(val)
	ch2 <- 20
}

func readChan(ch <-chan int) {
	// Invalid to write to chan
	val := <-ch
	fmt.Println(val)
	// ch <- 10
}

func writeChan(ch chan<- int) {
	// Invalid to read from chan
	// val := <-ch
	// fmt.Println(val)
	ch <- 10
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go writeChan(ch1)
	go defaultChan(ch1, ch2)
	go readChan(ch2)

	time.Sleep(1 * time.Second)
}
