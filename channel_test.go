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
	channel <- "hello world" // kirim data kedalam channel
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

// select channel
// digunakan ketika ingin menerima data dari beberapa channel
// urutan output mengikuti channel mana yang lebih dulu mengirim data
func TestSelectChannelLoop(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "data dari channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "data dari channel 2"
	}()

	// perulangan harus sesuai, kalau kelebihan bisa deadlock
	for i := 0; i < 2; i++ {
		select {
		case data := <-channel1:
			fmt.Println(data)
		case data := <-channel2:
			fmt.Println(data)
		}
	}
}

// select channel
// select disini akan memilih channel mana yg paling cepat, yg lain tidak di print
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Data dari Channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Data dari Channel 2"
	}()

	select {
	case data := <-channel1:
		fmt.Println(data)
	case data := <-channel2:
		fmt.Println(data)
	}
}

// select channel
// runtime go akan memilih salah satu channel yang siap secara acak,
// karena kedua channel siap pada waktu yang sama
func TestSelectRandom(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "Channel 1"
	}()

	go func() {
		channel2 <- "Channel 2"
	}()

	select {
	case data := <-channel1:
		fmt.Println(data)
	case data := <-channel2:
		fmt.Println(data)
	}
}

// default select
// perlu dibuat agar kalau ada channel yg tidak ada datanya, dia tidak deadlock
func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2: ", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
