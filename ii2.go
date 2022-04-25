// Blocking channel version
// Test 1
// 1
// 1 1000000
// Case #1: 525050
// 461.72ms
// Test 2
// 1
// 1 1000000000
// signal: killed

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	interesting := make(chan int)
	var count int
	var t int
	fmt.Scanf("%d", &t)

	for test := 1; test < t+1; test++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		start := time.Now()
		for num := a; num < b+1; num++ {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				s := 0
				p := 1
				integer := n
				var digit int
				for n > 0 {
					digit = n % 10
					n /= 10
					s += digit
					p *= digit
				}
				if p%s == 0 {
					interesting <- integer
				}
			}(num)
		}
		go func() {
			for {
				select {
				case <-interesting:
					count++
				}
			}
		}()
		wg.Wait()
		duration := time.Since(start)
		fmt.Printf("Case #%d: %d\n", test, count)
		fmt.Println(duration)
	}

}
