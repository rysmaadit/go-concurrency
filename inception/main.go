package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go func(ch chan int) {
		ch2 := make(chan int)
		go func(ch chan int) {
			ch3 := make(chan int)
			go func(ch chan int) {
				ch <- 3
			}(ch3)
			fmt.Println(<-ch3)
			ch <- 2
		}(ch2)
		fmt.Println(<-ch2)
		ch <- 1
	}(ch1)
	fmt.Println(<-ch1)
}
