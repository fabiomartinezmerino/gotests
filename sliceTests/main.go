package main

import (
	"fmt"
)

func main() {
	s1 := make([]string, 2, 8)
	s1[0] = "Hola"
	s1[1] = "Mundo"
	fmt.Printf("Capacidad: %v Longitud: %v Valor: %v", cap(s1), len(s1), s1)
}
