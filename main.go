package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan string, 2)
	go func() {
		time.Sleep(30 * time.Millisecond)
		out <- "slow"
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		out <- "fast"
	}()

	select {
	case v := <-out:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
