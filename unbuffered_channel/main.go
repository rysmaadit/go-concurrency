package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	message := make(chan string)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			message <- strconv.Itoa(i)
		}(i)
	}

	_, send := <-message

	if send {
		for elem := range message {
			fmt.Println(elem)
		}
	}

	wg.Wait()
	close(message)
}
