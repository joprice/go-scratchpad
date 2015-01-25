package main

// see https://mmcgrana.github.io/2012/09/go-by-example-timers-and-tickers.html

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("got tick")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
}
