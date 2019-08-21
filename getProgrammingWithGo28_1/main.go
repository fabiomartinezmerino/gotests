package main

import (
	"fmt"
	"os"
)

func main() {

	err := writeProverbs("goProverbs.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}

func writeProverbs(name string) error {

	f, err := os.Create(name)

	if err != nil {
		return err

	}
	defer f.Close()
	_, err = fmt.Fprintln(f, "Errors are values")

	if err != nil {
		//f.Close()
		return err

	}

	_, err = fmt.Fprintln(f, "Do not just check errors handle them gracefully")
	//f.Close()
	return err

}
