package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeProcess() chan int {
	r := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Intn(100)
	}()

	return r
}

func sumSquare(a, b int) int {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeProcess(), longTimeProcess()
	fmt.Println(sumSquare(<-a, <-b))
}
