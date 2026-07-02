package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
 * sync.Map mirip dengan map golang,
 * yang membedakan map ini aman untuk penggunaan concurrent menggunakan goroutine
 */

func AddtoMap(data *sync.Map, group *sync.WaitGroup, value int) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddtoMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
