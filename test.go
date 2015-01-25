package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func from(connection chan int) {
	connection <- rand.Intn(100)
}

func to(connection chan int) {
	i := <-connection
	fmt.Printf("got %d\n", i)
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	connection := make(chan int)
	go from(connection)
	go to(connection)
	time.Sleep(3000 * time.Millisecond)
}
