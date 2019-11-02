package main

import (
	"fmt"
	"log"
)

func main() {

	//slice to store integers

	sIntegers := make([]int, 10, 10)

	//ask for integers

	for counter := 0; counter < 10; counter++ {

		fmt.Printf("There remain %d integers to type - Please type another one\n", 10-counter)
		_, err := fmt.Scanf("%d", &sIntegers[counter])
		if err != nil {
			log.Fatal("Error - invalid integer - quitting")
		}

	}

	// test if slice is properly loaded
	for i, value := range sIntegers {

		fmt.Printf("Initial and NonSorted Position: %d - Value: %d\n", i, value)
	}

	// call BubbleSort to sort slice

	BubbleSort(sIntegers)

	fmt.Println("--- ---")

	// print slice sorted smaller numbers first

	for i, value := range sIntegers {

		fmt.Printf("Final and Sorted Position: %d - Sorted Value: %d\n", i, value)
	}

}

func BubbleSort(s []int) {

	for sweep_iteration := 0; sweep_iteration < len(s); sweep_iteration++ {
		/*For debugging to see how many times sweep must pass
		fmt.Printf("Sweep pass number: %d\n", sweep_iteration)
		*/
		// if an iteration sweep happens with no swapping it means entries are in proper place so no need toproceed with further iterations
		sorted := true
		//for each sweep we go through the slice and swap contiguous entries if they are not in correct place
		for counter := 0; counter < len(s)-1; counter++ {
			if s[counter] > s[counter+1] {
				swap(s, counter)
				sorted = false
			}
		}
		if sorted {
			break
		}

	}
}

func swap(s []int, index int) {
	s[index], s[index+1] = s[index+1], s[index]

}
