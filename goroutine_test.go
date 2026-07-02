package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // goroutine tidak bisa menangkap function yang memiliki return value
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display: ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := range 100000 {
		go DisplayNumber(i) // goroutine sangat ringan sekitar 2kb
	}

	time.Sleep(5 * time.Second)
}
