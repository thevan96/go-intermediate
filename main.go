package main

import (
	"fmt"
	"sync"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	chunkSize := (n + 3) / 4
	var total int
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if start >= n {
			start = n
		}

		if end > n {
			end = n
		}

		wg.Add(1)
		go func(i, j int) {
			defer wg.Done()
			partialSum := 0
			for _, value := range nums[i:j] {
				partialSum += value
			}
			mu.Lock()
			total += partialSum
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()
	fmt.Println(total)
}
