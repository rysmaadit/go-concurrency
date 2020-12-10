package main

import (
	"fmt"
	"time"
)

func process(ch <-chan int) {
	//ch <- 2 // this will break
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int, 2)

	// this kind of operation need 1 buffer
	ch <- 2
	ch <- 2

	go process(ch)
	go process(ch)

	//fmt.Println(<-ch)
	time.Sleep(time.Second * 3)
}
