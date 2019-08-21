package main

import (
	"fmt"
)

func main() {
	/*
		colors := map[string]string{
			"red":   "#FF0000",
			"green": "#008000",
		}
	*/
	/*
		var colors map[string]string
		colors = make(map[string]string)

		colors["red"] = "FF0000"
	*/

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
	}

	colors["white"] = "#FFFFFF"
	fmt.Printf("Map Name: %v\n", colors)
	/*delete(colors, "red")
	fmt.Printf("Map Name: %v\n", colors)
	*/
	iterateOverMap(colors)

	iterateOverMap(colors)
}

func iterateOverMap(m map[string]string) {

	for index, value := range m {
		fmt.Printf("Index: %v Value: %v\n", index, value)
	}

	m["white"] = "Blanco Modificao"

}
