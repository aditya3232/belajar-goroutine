package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// membaut channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // pastikan channel ditutup setelah digunakan

	/*
	 * jika hanya ada pengirim atau hanya ada penerima, goroutine akan terblock
	 * jika semua goroutine terblock, program akan deadlock
	 */
	go func() {
		time.Sleep(1 * time.Second)
		channel <- "data no. 1"
		fmt.Println("berhasil kirim data ke channel")

	}()

	// diterima oleh goroutine utama
	data := <-channel
	fmt.Println("result data from channel: ", data)

	time.Sleep(2 * time.Second)

}

// channel sebagai parameter
// parameter channel chan string, otmatis pass by reference (channel aslinya)
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "data no. 1" // kirim data kedalam channel
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel) // goroutine aka kirim channel ke sini

	data := <-channel // tinggal listen data channel
	fmt.Println("result data from channel: ", data)

	time.Sleep(2 * time.Second)
}
