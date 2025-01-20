package geometry

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %d", c.radius)
}

type Square struct {
	side int
}

func NewSquare(side int) *Square {
	return &Square{side: side}
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %d", s.side)
}
