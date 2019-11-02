package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		f float64

		s string
	)

	fmt.Println("Enter a floating point number")
	fmt.Scanf("%f", &f)
	s = strconv.FormatFloat(f, 'f', -1, 64)
	s = s[:strings.LastIndexByte(s, '.')]
	fmt.Printf("%s\n", s)

	/* simpler no need to import strconv and strings

	var (
		f float64
	)

	fmt.Println("Enter a floating point number")
	fmt.Scanf("%f", &f)
	fmt.Printf("%d", int64(f))

	*/
}
