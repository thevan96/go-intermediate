package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() (int, bool) {
	n := len(s.items)
	if n == 0 {
		return 0, false
	}

	top := s.items[n-1]
	s.items = s.items[:n-1]

	return top, true
}

func main() {
	var s Stack
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		s.Push(n)
	}

	for {
		num, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
