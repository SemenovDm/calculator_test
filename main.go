package main

import (
	"calculator_test/calculator"
	"calculator_test/shapes"
	"fmt"
)

func main() {

	circle1 := shapes.Circle{Radius: 5}
	circle2 := shapes.Circle{Radius: 10}

	rectangle1 := shapes.Rectangle{Width: 10, Heigh: 10}

	triangle1 := shapes.Triangle{Base: 2, Height: 5}

	circles := []shapes.Circle{circle1, circle2}
	rectanles := []shapes.Rectangle{rectangle1}
	triangles := []shapes.Triangle{triangle1}

	totalArea := calculator.TotalArea(circles, triangles, rectanles)

	fmt.Println(totalArea)
	fmt.Println(circle1)
	fmt.Println(circle2)

}
