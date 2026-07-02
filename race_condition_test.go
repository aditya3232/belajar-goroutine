package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// race condition
// terjadi ketika beberapa goroutine mengakses variable yg sama secara bersamaan
func TestRaceCondition(t *testing.T) {
	x := 0 // sharing variable

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter: ", x)
}
