package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)

	go func() {
		ch <- 1
	}()

	time.Sleep(time.Millisecond * 10)

	go func() {
		ch <- 2
	}()

	time.Sleep(time.Millisecond * 10)

	go func() {
		ch <- 3
	}()

	time.Sleep(time.Millisecond * 10)

	go func() {
		ch <- 4
	}()

	time.Sleep(time.Millisecond * 10)

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}
