// Concurrent brute-force with Mutex

// Test 1
// 1
// 1 1000000
// Case #1: 534358
// 210.738583ms

// Test 2
// 1
// 1 1000000000
// Case #1: 670349659
// 2m59.938570041s

package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	var t int
	fmt.Scanf("%d", &t)

	for test := 1; test < t+1; test++ {
		var interesting int
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
					m.Lock()
					interesting++
					m.Unlock()
				}
			}(num)
		}
		wg.Wait()
		fmt.Printf("Case #%d: %d\n", test, interesting)
	}

}
