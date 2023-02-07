package main

import "fmt"

func main() {

	var stack []string

	stack = append(stack, "singh")
	stack = append(stack, "Shishank")

	for len(stack) > 0 {
		// print top element
		n := len(stack) - 1
		fmt.Println(stack[n])

		//pop
		stack = stack[:n]
		fmt.Println(stack)
	}
}
