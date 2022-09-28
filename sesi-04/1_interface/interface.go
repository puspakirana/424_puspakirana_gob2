package main

import (
	"fmt"
	"math"
	"strings"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectagle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectagle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectagle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

// Type assertion
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectagle{width: 3, height: 2}

	//Type
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)
	fmt.Println(strings.Repeat("-", 25))

	//Area & Perimeter
	fmt.Println("Circle area", c1.area())
	fmt.Println("Circle perimeter", c1.perimeter())
	fmt.Println("Rectangle area", r1.area())
	fmt.Println("Rectangle perimeter", r1.perimeter())
	fmt.Println(strings.Repeat("-", 25))

	print("Rectangle", r1)
	print("Circle", c1)

	//Type Assertion
	//c1.volume() -> Salah
	//c1.(circle).volume() -> Benar

	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}
	fmt.Println(strings.Repeat("-", 25))

}