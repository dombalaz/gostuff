package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func putRandoms(ch chan int) {
	for range 10 {
		ch <- rand.Int()
	}
}

func launchCap() {
	ch := make(chan int, 20)
	done := make(chan struct{})

	go func() {
		putRandoms(ch)
	}()
	go func() {
		putRandoms(ch)
	}()

	go func() {
		defer close(done)
		for range cap(ch) {
			//fmt.Println(<-ch)
			_ = <-ch
		}
	}()
	<-done
}

func launchClose() {
	ch := make(chan int)
	var wg sync.WaitGroup
	var wgDone sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		putRandoms(ch)
	}()
	go func() {
		defer wg.Done()
		putRandoms(ch)
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	wgDone.Add(1)
	go func() {
		defer wgDone.Done()
		for v := range ch {
			//fmt.Println(v)
			_ = v
		}
	}()
	wgDone.Wait()
}

func main() {
	launchClose()
	fmt.Println("===")
	launchCap()
}
