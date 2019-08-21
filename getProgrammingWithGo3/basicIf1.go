/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var number1, number2 int

	number1 = rand.Intn(100)
	number2 = rand.Intn(100)

	if number1 > number2 {
		fmt.Printf("El numero mayor es:%d\n", number1)
	} else if number2 > number1 {
		fmt.Printf("El numero mayor es:%d\n", number2)
	} else {
		fmt.Printf("Los dos números %d %d son iguales", number1, number2)
	}

}
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence, command string = "I'm going into the cave", "out"
	if strings.Contains(sentence, command) {
		fmt.Println("I'm going out")

	} else if strings.Contains(sentence, "in") {
		fmt.Println("I will stay")
	} else {
		fmt.Println("I do not know")
	}

}
