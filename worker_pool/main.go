package main

import (
	"log"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	for w := 1; w <= 10; w++ {
		jobs <- w
	}

	close(jobs)

	for w := 1; w <= 10; w++ {
		<-results
	}
}

func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		log.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		log.Println("worker", id, "ended job", j)
		result <- j * 2
	}
}
