package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// untuk menunggu beberapa goroutine selesai dieksekusi, sebelum melanjutkan ke proses selanjutnya
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // jika lupa defer Done, bisa Deadlock, menunggu terus
	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("complete")
}
