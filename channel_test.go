package belajar_goroutine

import (
	"fmt"
	"strconv"
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

// channel in out (penanda channel untuk in dan out di parameter)
func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "data no. 1"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("result data from channel: ", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

// buffered channel
// cocok digunakan ketika pengirim lebih cepat daripada penerima
// dia adalah kapasitas antrian dalam buffer
// dengan adanya buffer, data yang dimasukkan akan ditampung di antrian,
// sehigga tidak deadlock, kalau belum ada penerima,
// tapi kalau buffer nya full, dan tidak ada yg ambil, ya bakalan deadlock juga kalau ada data baru
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "adit"
		channel <- "ichsan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("selesai")
	time.Sleep(1 * time.Second)
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		defer close(channel) // agar tidak deadlock close message
		for i := range 10 {
			data := "data no. " + strconv.Itoa(i)
			channel <- data
		}
	}()

	for data := range channel {
		fmt.Println("menerima ", data)
	}

	fmt.Println("selesai")
}
