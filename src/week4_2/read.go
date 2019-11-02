package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//Customized type to store first name and second name read from each file register
type person struct {
	fname, sname string
}

//Customized type to store several person in a slice type fashion
type slpeople []person

func main() {

	// Prompt for file name - I suggest you type people.txt - 'x' or 'X' to exit

	var n string //variable to store filename to read from

	for {

		fmt.Println("Please input filename, no spaces please, x or X if you wish to quit: ")

		fmt.Scanf("%s", &n)
		// now we try to get a handle on the file, if it doesn't exist or another problem should occur we simply try again
		// poor error handling we hone this skill throughout the course
		if strings.ToLower(n) == "x" {
			fmt.Println("Quitting...")
			time.Sleep(time.Second * 2)
			os.Exit(0)

		}

		if fileExists(n) {

			break

		} else {
			fmt.Println("No file with this name is found....")
			continue
		}
	}

	// Open file, manage naively possible errors

	f, err := os.Open(n)

	if err != nil {
		log.Printf("An error has happened opening file: %v", err)
		os.Exit(1) // try again in case filename is incorrect

	}

	// closing file drudgery - using defer to make sure this code executes before exiting

	defer func() {
		if err := f.Close(); err != nil {
			// this time if something goes wrong simply exit with 1 return code
			log.Fatal(err)
		}
	}() // invoking anonymous funtion

	// reading file one line at a time and store results

	psl := make(slpeople, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		reg := scanner.Text()
		//create a person instance and append it to slpeople slice
		buff := strings.Fields(reg)
		// cut down to 20 characters first and second name before stow them away
		if len(buff[0]) > 19 {
			buff[0] = buff[0][0:19]
		}
		if len(buff[1]) > 19 {
			buff[1] = buff[1][0:19]
		}
		// put the struct instance into a slice
		psl = append(psl, (person{fname: buff[0], sname: buff[1]}))
		//fmt.Printf("Register: %v\n", strings.Fields(reg))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Process slice of structs

	for index, person_instance := range psl {

		fmt.Printf("Person number: %d - Name:%s - Second Name:%s\n", index, person_instance.fname, person_instance.sname)
	}

}

// function to test if file exists

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
