package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(encode([]string{"lint", "code", "love", "you"}))
}

func encode(str []string) string {
	var sb strings.Builder
	for _, v := range str {
		c := v + "\\"
		sb.WriteString(c)
	}
	return sb.String()
}

// func decode(str string) []string {

// }
