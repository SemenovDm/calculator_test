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
	square1 := shapes.Square{Length: 4}
	trapezoid1 := shapes.Trapezoid{Length1: 2, Length2: 5, Base: 5}

	circles := []shapes.Circle{circle1, circle2}
	rectanles := []shapes.Rectangle{rectangle1}
	triangles := []shapes.Triangle{triangle1}
	squares := []shapes.Square{square1}
	trapezoids := []shapes.Trapezoid{trapezoid1}

	totalArea := calculator.TotalArea(circles, triangles, rectanles, squares, trapezoids)

	fmt.Println(totalArea)

	var shapeName string
	// var Width float64
	// var Height float64
	var Length float64
	// var Length1 float64
	// var Length2 float64
	// var Base float64
	// var Radius float64

	fmt.Println("Введите название фигуры")
	fmt.Scan(&shapeName)

	if shapeName == "Square" {
		fmt.Println("Введите длину стороны квадрата")
		fmt.Scan(&Length)
		sq := shapes.Square{Length: Length}
		fmt.Printf("Площадь квадрата: %.2f", sq.Area())
	}

}
