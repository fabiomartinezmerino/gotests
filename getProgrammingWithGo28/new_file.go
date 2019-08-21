package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	dirs, err := ioutil.ReadDir("/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Printf("Dir name: %v\n", dir.Name())
		} else {
			fmt.Printf("File name: %v\n", dir.Name())
		}
	}

}
