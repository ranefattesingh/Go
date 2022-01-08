package main

import "fmt"

type Shape struct {
	name string
	side float64
}

type Rectangle struct {
	Shape
	newSide float64
}

func (r Rectangle) CalculateArea() float64 {
	return 0.5 * r.side * r.newSide
}

func (r Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.side + r.newSide)
}

func (r Rectangle) GetName() string {
	return r.name
}

type Square struct {
	Shape
}

func (s Square) CalculateArea() float64 {
	return s.side * s.side
}

func (s Square) CalculatePerimeter() float64 {
	return 4 * s.side
}

func (s Square) GetName() string {
	return s.name
}

type Circle struct {
	Shape
}

func (c Circle) CalculateArea() float64 {
	return 3.14 * c.side * c.side
}

func (c Circle) CalculatePerimeter() float64 {
	return 2 * 3.14 * c.side
}

func (c Circle) GetName() string {
	return c.name
}

type Interface interface {
	CalculateArea() float64
	CalculatePerimeter() float64
	GetName() string
}

type Larger struct {
	Interface
}

func (c Larger) GetLargeArea() float64 {
	return c.CalculateArea()
}

func (c Larger) GetLargePerimeter() float64 {
	return c.CalculatePerimeter()
}

func (c Larger) GetName() string {
	return c.Interface.GetName()
}

// Simple printing functions
func PrintArea(shape Interface) {
	fmt.Println(shape.GetName(), shape.CalculateArea())
}

func PrintPerimeter(shape Interface) {
	fmt.Println(shape.GetName(), shape.CalculatePerimeter())
}

func WhoIsLarger(i1, i2, i3 Interface) Larger {
	if i1.CalculateArea() > i2.CalculateArea() && i1.CalculateArea() > i3.CalculateArea() {
		return Larger{i1}
	} else if i2.CalculateArea() > i3.CalculateArea() {
		return Larger{i2}
	} else {
		return Larger{i3}
	}
}

func main() {
	square := Square{Shape: Shape{side: 10, name: "Square"}}
	rectangle := Rectangle{Shape: Shape{side: 10, name: "Rectangle"}, newSide: 20}
	circle := Circle{Shape: Shape{side: 10, name: "Circle"}}

	PrintArea(square)
	PrintPerimeter(square)

	PrintArea(rectangle)
	PrintPerimeter(rectangle)

	PrintArea(circle)
	PrintPerimeter(circle)

	fmt.Println("Larger Shape")
	largerShape := WhoIsLarger(square, rectangle, circle)
	PrintArea(largerShape)
	PrintPerimeter(largerShape)

}
