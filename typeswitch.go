package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Double %v", v)
	case string:
		fmt.Printf("String %v", v)
	default:
		fmt.Printf("Unknown type %v", v)
	}
}
func main() {
	do(22)
	do("hello")
	do(22.4)
}
