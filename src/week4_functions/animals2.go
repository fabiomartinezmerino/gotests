package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define cow "object" that implements Animal interface
type cow struct {
	animal_type, food_eaten, locomotion_method, spoken_language string
}

func (c cow) Eat() {
	fmt.Printf("%v eats %v\n", c.animal_type, c.food_eaten)
}

func (c cow) Move() {
	fmt.Printf("%v locomotion method is: %v\n", c.animal_type, c.locomotion_method)
}

func (c cow) Speak() {
	fmt.Printf("%v sound is: %v\n", c.animal_type, c.spoken_language)
}

// Define bird "object" that implements Animal interface
type bird struct {
	animal_type, food_eaten, locomotion_method, spoken_language string
}

func (b bird) Eat() {
	fmt.Printf("%v eats %v\n", b.animal_type, b.food_eaten)
}

func (b bird) Move() {
	fmt.Printf("%v locomotion method is: %v\n", b.animal_type, b.locomotion_method)
}

func (b bird) Speak() {
	fmt.Printf("%v sound is: %v\n", b.animal_type, b.spoken_language)
}

// Define snake "object" that implements Animal interface
type snake struct {
	animal_type, food_eaten, locomotion_method, spoken_language string
}

func (s snake) Eat() {
	fmt.Printf("%v eats %v\n", s.animal_type, s.food_eaten)
}

func (s snake) Move() {
	fmt.Printf("%v locomotion method is: %v\n", s.animal_type, s.locomotion_method)
}

func (s snake) Speak() {
	fmt.Printf("%v sound is: %v\n", s.animal_type, s.spoken_language)
}

func main() {

	// Istantiate container for Animals
	c := make(map[string]Animal)

	// Prompt for command string

	for {
		fmt.Println("Enter command sequence: >")

		reader := bufio.NewReader(os.Stdin)
		//Read from console and store into string
		str, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		strSl := strings.Fields(str)
		// strSl should have 3 entries, if this is not the case dismiss command and exit
		if len(strSl) != 3 {
			log.Printf("Incorrect number of Arguments - Program will exit")
			continue
		}
		// Process command Type - newanimal - query
		switch commandType := strings.ToLower(strSl[0]); {
		case commandType == "newanimal": // new Animal creation is requested
			switch animalType := strings.ToLower(strSl[2]); {
			case animalType == "cow":
				//Instantiate a cow - name is at 2nd entry of slice - and store it
				c[strSl[1]] = cow{animal_type: "cow", food_eaten: "grass", locomotion_method: "walk", spoken_language: "moo"}
			case animalType == "bird":
				//Instantiate a cow - name is at 2nd entry of slice - and store it
				c[strSl[1]] = bird{animal_type: "bird", food_eaten: "worms", locomotion_method: "fly", spoken_language: "peep"}
			case animalType == "snake":
				//Instantiate a cow - name is at 2nd entry of slice - and store it
				c[strSl[1]] = snake{animal_type: "snake", food_eaten: "mice", locomotion_method: "slither", spoken_language: "hsss"}
			default:
				//Not known animal type - cannot be processed
				log.Printf("You have entered a not known species - please try again")
				continue
			}
			fmt.Printf("Created it!\n")
		case commandType == "query": // information about a stored Animal is requested
			//Retrieve Animal - name can be located on 1 position
			anm := c[strSl[1]]

			if anm != nil {
				fmt.Printf("You have requested information on %v \n", strSl[1])
				switch queryType := strings.ToLower(strSl[2]); {
				case queryType == "eat":
					anm.Eat()
				case queryType == "move":
					anm.Move()
				case queryType == "speak":
					anm.Speak()
				default:
					log.Printf("You have asked for unsopported piece of information - please try again")
					continue
				}
			} else {
				log.Printf("Problem retrieving animal - please try again")
				continue
			}

		default:
			log.Printf("Command not supported - plaease try again")
			continue
		}
	}
}
