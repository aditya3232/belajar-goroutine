package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// digunakan untuk menyimpan dan menggunakan kembali (reuse) objek yang sering dibuat dan dibuang,
// sehingga mengurangi alokasi memori dan beban garbage collector
func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("data 1")
	pool.Put("data 2")
	pool.Put("data 3")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}
