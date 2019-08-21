package main

import (
	"fmt"
)

func main() {

	var dividend, divisor int = 100, 47
	//dividend = 100
	//divisor = 47
	fmt.Printf("Result: %d: \n", dividend%divisor)

	var firstbyte byte = 3
	var movement uint = 1
	fmt.Printf("Original number: %b \n", firstbyte)
	fmt.Printf("Result: %b \n", firstbyte<<movement)

}
