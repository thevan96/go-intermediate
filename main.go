package main

import (
	"errors"
	"fmt"
)

func safeDivide(a, b int) (q int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("divide by zero")
		}
	}()

	q = a / b
	return q, nil
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if div, err := safeDivide(a, b); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("result: %d\n", div)
	}
}
