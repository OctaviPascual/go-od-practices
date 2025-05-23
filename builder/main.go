package main

import (
	"math/rand"
	"testing"
	"time"
)

type Car struct {
	createdAt  time.Time
	color      string
	hasSunroof bool
}

func NewCar(createdAt time.Time, color string, hasSunroof bool) *Car {
	return &Car{
		createdAt:  createdAt,
		color:      color,
		hasSunroof: hasSunroof,
	}
}

type testingCarBuilder struct {
	*Car
}

func NewCarForTesting(t testing.TB) *testingCarBuilder {
	t.Helper()
	return &testingCarBuilder{
		&Car{
			createdAt:  time.Now(),
			color:      randString(6),
			hasSunroof: false,
		},
	}
}

func (t *testingCarBuilder) WithCreatedAt(createdAt time.Time) *testingCarBuilder {
	t.createdAt = createdAt
	return t
}

func (t *testingCarBuilder) WithColor(color string) *testingCarBuilder {
	t.color = color
	return t
}

func (t *testingCarBuilder) WithSunroof() *testingCarBuilder {
	t.hasSunroof = true
	return t
}

func (t *testingCarBuilder) Build() *Car {
	return t.Car
}

func randString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
