package belajar_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// sync atomic
// package yg digunakan untuk menggunakan data primitive secara aman pada proses concurrent
// jadi gk perlu mutex untuk locking data primitive
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 1; j <= 100; j++ {
				// x = x + 1
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()

	fmt.Println("counter: ", x)
}
