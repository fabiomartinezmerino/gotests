package main

import (
	"fmt"
	"math"
)

func main() {

	s := newSquare()
	t := newTriangle()
	fmt.Printf("Triangulo: %+v", t)
	printArea(s)
	printArea(t)

}

type square struct {
	sideLength float64
}

type triangle struct {
	sideLength float64
	height     float64
}

func newSquare() square {

	var lado float64
	fmt.Println("Creación Cuadrado")
	fmt.Println("Introduce longitud de lado")
	fmt.Scanf("%f", &lado)
	return square{sideLength: lado}

}

func (s square) getArea() float64 {

	fmt.Println("Imprimiendo area cuadrado")
	fmt.Printf("Longitud de lado: %v\n", s.sideLength)
	return math.Pow(s.sideLength, 2)

}

func newTriangle() triangle {
	var base, altura float64
	fmt.Println("Creación Triángulo\n")
	fmt.Println("Introduce longitud de base\n")
	fmt.Scanf("%f", &base)
	fmt.Println("Introduce longitud de altura\n")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Base: %f, Altura: %f\n", base, altura)
	return triangle{sideLength: base, height: altura}

}

func (t triangle) getArea() float64 {
	fmt.Println("Imprimiendo area triángulo")
	fmt.Printf("Altura: %v Base: %v\n", t.height, t.sideLength)
	return (0.5 * t.height * t.sideLength)
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {

	fmt.Printf("El area del %T es: %f cm\n", s, s.getArea())

}
