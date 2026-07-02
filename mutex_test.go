package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync mutex (mutual exclusion)
// digunakan untuk lock variable yg diakses goroutine
// seperti membuat antrian kepada akses variable yg dilakukan goroutine
func TestMutex(t *testing.T) {
	var mu sync.Mutex
	x := 0 // sharing variable

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mu.Lock()
				x = x + 1
				mu.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("counter: ", x)
}
