package main

import (
	"fmt"
)

func Sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(Sum(5, 10))
	fmt.Println(sub(1, 4))
}
