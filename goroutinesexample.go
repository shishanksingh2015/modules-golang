package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int, c chan int) {
	defer wg.Done()
	for i := range c {
		fmt.Printf("Worker with ID %d finished with value %d", id, i)
	}
}

func main() {

}
