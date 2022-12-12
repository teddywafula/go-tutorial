package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) area2() int {
	return r.height * r.width
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println("go methods using struct")
	r := rect{width: 20, height: 10}
	fmt.Println(r)
	fmt.Println("area of rectangle is :", r.area())
	fmt.Println("perimeter of rectangle is :", r.perimeter())
	fmt.Println("Area of rectangle with pointer is :", r.area2())
}
