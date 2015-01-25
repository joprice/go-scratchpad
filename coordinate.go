package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, start chan bool, w *sync.WaitGroup) {
	defer w.Done()
	msg := <-start
	fmt.Println("running", i, msg)
	time.Sleep(1 * time.Second)
}

func main() {
	var w sync.WaitGroup
	count := 100
	w.Add(count)

	start := make(chan bool)

	for i := 0; i < count; i++ {
		go worker(i, start, &w)
	}

	fmt.Println("start")
	// closing a channel is considered sending it a message
	close(start)
	// workers running
	w.Wait()
	fmt.Println("stop")
}
