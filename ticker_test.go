package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for {
		select {
		case currentTime := <-ticker.C:
			fmt.Println(currentTime)

		case <-done:
			fmt.Println("ticker stopped")
			return
		}
	}
}

// Ticker digunakan untuk menjalankan suatu pekerjaan secara berkala dengan interval waktu tertentu.
// bisa digunakan untuk auto refresh data
// health check
// mengirim metrics
// atau sinkronisasi data
func TestTick(t *testing.T) {
	tick := time.Tick(time.Second)
	stop := time.After(5 * time.Second)

	for {
		select {
		case currentTime := <-tick:
			fmt.Println(currentTime)

		case <-stop:
			fmt.Println("stop")
			return
		}
	}
}
