package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var (
		i  int
		s  string
		sl []int
	)

	sl = make([]int, 0, 10)

	for {
		fmt.Println("Please, enter an integer, X to quit: ")
		fmt.Scanf("%s", &s)
		if strings.ToUpper(s) == "X" {
			break
		} else {
			i, _ = strconv.Atoi(s)
			sl = append(sl, i) //append takes charge of increasing slice capacity every time it is needed
			sort.Ints(sl)
			fmt.Printf("Slice elements in sorted order: %v\n", sl)
			fmt.Println("---- Now ordered using a for loop ----")
			for index, value := range sl {
				fmt.Printf("Index: %d Value: %d\n", index, value)
			}
			fmt.Println("---- Length and Capacity ----")
			fmt.Printf("Length: %d Capacity: %d\n", len(sl), cap(sl))

		}

	}

	fmt.Println("Program will exit...")

}
