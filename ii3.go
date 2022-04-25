// Concurrent brute-force with atomic add

// Test 1
// 1
// 1 1000000
// Case #1: 534358
// 202.710917ms

// Test 2
// 1
// 1 1000000000
// Case #1: 670349659
// 2m49.322227875s

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var t int
	fmt.Scanf("%d", &t)

	for test := 1; test < t+1; test++ {
		var interesting int32
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		for num := a; num < b+1; num++ {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				s := 0
				p := 1
				var digit int
				for n > 0 {
					digit = n % 10
					n /= 10
					s += digit
					p *= digit
				}
				if p%s == 0 {
					atomic.AddInt32(&interesting, 1)
				}
			}(num)
		}
		wg.Wait()
		fmt.Printf("Case #%d: %d\n", test, interesting)
	}
}
