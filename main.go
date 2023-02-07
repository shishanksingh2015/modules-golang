package main

import "fmt"

func main() {
	p := new(chan int)
	c := make(chan int)

	fmt.Println(p)
	fmt.Println(c)
}
