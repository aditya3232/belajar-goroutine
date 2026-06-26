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
