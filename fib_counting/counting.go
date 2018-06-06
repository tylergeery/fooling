package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/tylergeery/fooling/fib"
)

func count(ch chan int) {
	inc := 0

	for {
		select {
		case <-time.After(1 * time.Second):
			inc++
			ch <- inc
		case _ = <-ch:
			return
		}
	}
}

func startFib(ch chan *big.Int) {
	ch <- fib.Fib_tail(166660, big.NewInt(1), big.NewInt(0))
}

func main() {
	ch_count := make(chan int)
	ch_fib := make(chan *big.Int)

	go count(ch_count)
	go startFib(ch_fib)

	for {
		select {
		case done := <-ch_fib:
			fmt.Printf("Fib(166660) = %s\n", done)
			ch_count <- 1
			return
		case count := <-ch_count:
			fmt.Println(count)
		}
	}

	go fmt.Println()
	<-ch_fib
}
