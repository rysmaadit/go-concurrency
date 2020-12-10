package main

import (
	"fmt"
	"time"
)

// this kind of channel as function argument that possibly could received and send value

func main() {
	ch := make(chan int, 2)
	go process(ch)

	// add time to give goroutines to run
	time.Sleep(time.Second * 1)
	fmt.Println(<-ch)
}

func process(ch chan int) {
	// value assignment on the given channel
	ch <- 2
	s := <-ch
	fmt.Println(s)
}
