package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func f(v *uint32, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		atomic.AddUint32(v, 1)
	}

	wg.Done()
}

func main() {
	v := uint32(1)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go f(&v, &wg)
	go f(&v, &wg)

	wg.Wait()

	fmt.Println(v)
}
