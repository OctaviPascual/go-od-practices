package main

import (
	"fmt"

	"github.com/OctaviPascual/go-testing-practices/interface_embedding/geometry"
)

// Useful blog about interface embedding:
// https://eli.thegreenplace.net/2020/embedding-in-go-part-2-interfaces-in-interfaces/

type ExtendedShape interface {
	geometry.Shape
	ExtendedRender() string
}

type extendedShapeImpl struct {
	geometry.Shape
}

func (s *extendedShapeImpl) ExtendedRender() string {
	return fmt.Sprintf("Extended: %s", s.Render())
}

func main() {
	var extendedShape ExtendedShape
	extendedShape = &extendedShapeImpl{geometry.NewCircle(2)}
	fmt.Println(extendedShape.Render())
	fmt.Println(extendedShape.ExtendedRender())

	extendedShape = &extendedShapeImpl{geometry.NewSquare(5)}
	fmt.Println(extendedShape.Render())
	fmt.Println(extendedShape.ExtendedRender())
}
