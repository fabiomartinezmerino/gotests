package main

import (
	"fmt"
)

func main() {
	var variable int

	for {

		fmt.Println("Introduce un numero:")
		fmt.Scanf("%d", &variable)
		if variable == 0 {
			break
		}
		switch variable {

		case 1:
			fmt.Println("Case 1")

		case 2:
			fmt.Println("Case 2")

		case 3:
			fmt.Println("Case 3")
		default:
			fmt.Println("None of the above")

		}

	}
}
