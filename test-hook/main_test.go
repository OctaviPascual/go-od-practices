package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testHookGreet = func(s string) string {
		return fmt.Sprintf("Goodbye, %s!", s)
	}
	t.Cleanup(func() { testHookGreet = nil })

	assert.Equal(t, "Goodbye, world!", greet("world"))
}
