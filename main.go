package main

import (
	"bufio"
	"fmt"
	"os"
)

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := make(chan int)
	go func() {
		defer close(a)
		for {
			var n int
			if _, err := fmt.Fscan(reader, &n); err != nil {
				break
			}
			a <- n
		}
	}()

	b := square(a)
	var sum int
	for n := range b {
		sum += n
	}
	fmt.Println(sum)
}
