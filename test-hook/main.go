package main

import (
	"fmt"
)

// Make sure this is private so users outside the package can't see it.
var testHookGreet func(string) string

func greet(s string) string {
	if h := testHookGreet; h != nil {
		return h(s)
	}
	return fmt.Sprintf("Hello, %s!", s)
}

func main() {
	fmt.Println(greet("world"))
}
