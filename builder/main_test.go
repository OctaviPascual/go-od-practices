package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("using newCar requires setting all fields even if they aren't needed", func(t *testing.T) {
		createdAt := time.Now()
		car := NewCar(createdAt, "red", true)

		assert.Equal(t, createdAt, car.createdAt)
	})

	t.Run("using newCarForTesting to set only the relevant fields for this test case", func(t *testing.T) {
		createdAt := time.Now()
		car := NewCarForTesting(t).WithCreatedAt(createdAt).Build()

		assert.Equal(t, createdAt, car.createdAt)
	})
}
