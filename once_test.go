package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(OnlyOnce) // memastikan fungsi yang berjalan hanya 1 kali, gorouitne pertama yang eksekusi
		}()
	}

	group.Wait()

	fmt.Println("counter: ", counter)
}
