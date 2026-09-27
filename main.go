package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kind := scanner.Text()

	scanner.Scan()
	dim, _ := strconv.ParseFloat(scanner.Text(), 64)

	var s Shape
	if kind == "circle" {
		s = Circle{Radius: dim}
	} else if kind == "square" {
		s = Square{Side: dim}
	}

	if s != nil {
		fmt.Printf("%.2f\n", s.Area())
	}
}
