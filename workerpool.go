package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for i := range jobs {
		fmt.Println("worker", id, "started job", i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", i)
		results <- i * 2
	}

}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= 10; a++ {
		fmt.Println(<-results)
	}
}
