package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func queryAll() int {
	ch := make(chan int)
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	return <-ch
}

func main() {
	for i := 0; i < 10; i++ {
		queryAll()
		runtime.GC()
		runtime.GC()
		runtime.GC()
		runtime.GC()
		runtime.GC()
		fmt.Printf("%d, #goroutines: %d\n", i, runtime.NumGoroutine())
	}
	runtime.GC()
	runtime.GC()
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
}
