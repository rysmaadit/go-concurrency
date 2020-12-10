package main

import "fmt"

// this kind of channel as function argument that only possible to send value

func process(ch chan<- int) {
	ch <- 2
}

func main() {
	ch := make(chan int, 2)
	go process(ch)
	go process(ch)

	// when receiving value from channel it deduct the channel buffers by one
	// buffer only reduced by reading a value from channel
	// value assignment from a channel do not reduce channel buffers

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
