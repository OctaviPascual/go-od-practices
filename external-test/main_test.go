// Add _test to the package name so test becomes external.
package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/OctaviPascual/go-od-practices/external-test"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, main.Sum(5, 5))
}

func TestSub(t *testing.T) {
	assert.Equal(t, 5, main.Sub(10, 5))
}
