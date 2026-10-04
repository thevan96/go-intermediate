package main

import (
	"fmt"
	"strconv"
)

type Logger struct {
}

func (l Logger) Log(msg string) {
	fmt.Printf("[log] %s\n", msg)
}

type Counter struct {
	Logger
	count int
}

func (c *Counter) Inc() {
	c.count++

	numRaw := strconv.Itoa(c.count)
	c.Log(numRaw)
}

func main() {
	var num int
	fmt.Scan(&num)
	var c Counter
	for i := 0; i < num; i++ {
		c.Inc()
	}
}
