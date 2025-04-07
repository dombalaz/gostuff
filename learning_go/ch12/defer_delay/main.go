package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer fmt.Println("1")
		defer time.Sleep(1 * time.Second)
		defer fmt.Println("2")
		// ^^^ These are not guaranteed to be executed!
		defer wg.Done()
		defer fmt.Println("3")
		defer time.Sleep(1 * time.Second)
		defer fmt.Println("4")
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()
	inMain := 2
	var fromGoroutine int
	count := 0
	for count < 2 {
		select {
		case ch2 <- inMain:
			count++
		case fromGoroutine = <-ch1:
			count++
		}
	}
	fmt.Println("main:", inMain, fromGoroutine)
	wg.Wait()
}
