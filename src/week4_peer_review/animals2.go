// Write a program which allows the user to create a set of animals and to get information about those animals.
// Each animal has a name and can be either a cow, bird, or snake.
// With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
// Each animal has a unique name, defined by the user.
// Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
// The following table contains the three types of animals and their associated data.

// Animal	Food eaten	Locomtion method	Spoken sound
// cow		grass		walk				moo
// bird		worms		fly					peep
// snake	mice		slither				hsss

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
// Your program should continue in this loop forever.
// Every command from the user must be either a “newanimal” command or a “query” command.

// Each “newanimal” command must be a single line containing three strings.
// The first string is “newanimal”.
// The second string is an arbitrary string which will be the name of the new animal.
// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

// Each “query” command must be a single line containing 3 strings.
// The first string is “query”.
// The second string is the name of the animal.
// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
// Your program should process each query command by printing out the requested data.

// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
// Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
// When the user creates an animal, create an object of the appropriate type.
// Your program should call the appropriate method when the user issues a query command.

package main

import (
	"fmt"
)

// Animal interface to bring together all animal structs
type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

// The Cow struct
type Cow struct {
	Name       string
	Food       string
	Locomotion string
	Noise      string
}

// getName function for the Cow struct. Returns the Cow's Name
func (cow Cow) getName() string {
	return cow.Name
}

// Eat function for the Cow struct. Print's the Cow's Food
func (cow Cow) Eat() {
	fmt.Printf("%v\n", cow.Food)
}

// Move function for the Cow struct. Prints the Cow's Locomotion
func (cow Cow) Move() {
	fmt.Printf("%v\n", cow.Locomotion)
}

// Speak function for the Cow struct. Prints the Cow's Noise
func (cow Cow) Speak() {
	fmt.Printf("%v\n", cow.Noise)
}

// The Bird struct, and corresponding functions Eat(), Move() & Speak()
type Bird struct {
	Name       string
	Food       string
	Locomotion string
	Noise      string
}

// getName function for the Bird struct. Returns the Bird's Name
func (bird Bird) getName() string {
	return bird.Name
}

// Eat function for the Bird struct. Prints the Bird's Food
func (bird Bird) Eat() {
	fmt.Printf("%v\n", bird.Food)
}

// Move function for the Bird struct. Prints the Bird's Locomotion
func (bird Bird) Move() {
	fmt.Printf("%v\n", bird.Locomotion)
}

// Speak function for the Bird struct. Prints the Bird's Noise
func (bird Bird) Speak() {
	fmt.Printf("%v\n", bird.Noise)
}

// The Snake struct, and corresponding functions Eat(), Move() & Speak()
type Snake struct {
	Name       string
	Food       string
	Locomotion string
	Noise      string
}

// getName function for the Snake struct. Returns the Snake's Name
func (snake Snake) getName() string {
	return snake.Name
}

// Eat function for the Snake struct. Prints the Snake's Food
func (snake Snake) Eat() {
	fmt.Printf("%v\n", snake.Food)
}

// Move function for the Snake struct. Prints the Snake's Locomotion
func (snake Snake) Move() {
	fmt.Printf("%v\n", snake.Locomotion)
}

// Speak function for the Snake struct. Prints the Snake's Noise
func (snake Snake) Speak() {
	fmt.Printf("%v\n", snake.Noise)
}

// checkCommandType checks if commandType is valid
func checkCommandType(commandType string) bool {
	if commandType == "newanimal" || commandType == "query" {
		return true
	} else {
		fmt.Printf("ERROR: Incorrect command type. Expected [newanimal | query], got: %v\n", commandType)
		return false
	}
}

// checkparam checks if the value in param is valid for the given commandType
func checkparam(commandType string, param string) bool {
	if commandType == "newanimal" {
		newSlice := []string{"cow", "bird", "snake"}
		if stringInSlice(param, newSlice) {
			return true
		} else {
			return false
		}
	} else if commandType == "query" {
		querySlice := []string{"eat", "move", "speak"}
		if stringInSlice(param, querySlice) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// stringInSlice checks if a string exists in a slice of strings (similar to Python's if x in [a,b,c])
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
}

// checkList iterates through animalList until the Name matches name. Then calls getInfo()
func checkList(animalList []Animal, name string, param string) {
	for _, i := range animalList {
		animalName := i.getName()
		if animalName == name {
			getInfo(i, param)
		}
	}
}

// getInfo calls the appropriate method for animal, depending on the param given
func getInfo(animal Animal, param string) {
	if param == "eat" {
		animal.Eat()
	} else if param == "move" {
		animal.Move()
	} else if param == "speak" {
		animal.Speak()
	}
}

func main() {
	var commandType string
	var animalName string
	var param string
	var animal Animal
	var animalList []Animal

	for {
		fmt.Printf("> ")
		fmt.Scan(&commandType, &animalName, &param)
		if checkCommandType(commandType) {
			if checkparam(commandType, param) {
				if commandType == "newanimal" {

					if param == "cow" {
						cow1 := Cow{animalName, "grass", "walk", "moo"}
						animal = cow1
						animalList = append(animalList, animal)
						fmt.Printf("Created it!\n")
					} else if param == "bird" {
						bird1 := Bird{animalName, "worms", "fly", "peep"}
						animal = bird1
						animalList = append(animalList, animal)
						fmt.Printf("Created it!\n")
					} else if param == "snake" {
						snake1 := Snake{animalName, "mice", "slither", "hsss"}
						animal = snake1
						animalList = append(animalList, animal)
						fmt.Printf("Created it!\n")
					}

				} else if commandType == "query" {
					checkList(animalList, animalName, param)
				}
			} else {
				fmt.Printf("ERROR: Incorrect animal type. Expecting [cow, bird, snake], received: %v\n", param)
			}
		}
	}
}
