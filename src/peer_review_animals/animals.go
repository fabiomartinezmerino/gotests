package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForAnimal() ([]string, error) {
	fmt.Printf("> ")
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	words := strings.Fields(line)
	return words, err
}

type Animal struct {
	typeOf     string
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Initialize(typeOf string) {
	animal.typeOf = typeOf
}

func (animal *Animal) Eat() {
	switch animal.typeOf {
	case "cow":
		animal.food = "grass"
	case "bird":
		animal.food = "worms"
	case "snake":
		animal.food = "mice"
	default:
		animal.food = "not supported yet"
	}
}

func (animal *Animal) Move() {
	switch animal.typeOf {
	case "cow":
		animal.locomotion = "walk"
	case "bird":
		animal.locomotion = "fly"
	case "snake":
		animal.locomotion = "slither"
	default:
		animal.locomotion = "not supported yet"
	}
}

func (animal *Animal) Speak() {
	switch animal.typeOf {
	case "cow":
		animal.noise = "moo"
	case "bird":
		animal.noise = "peep"
	case "snake":
		animal.noise = "hsss"
	default:
		animal.noise = "not supported yet"
	}
}

func main() {
	for {
		words, error := AskForAnimal()

		if error == nil {
			animalName := words[0]
			action := words[1]

			var animal Animal
			animal.Initialize(animalName)

			switch action {
			case "eat":
				animal.Eat()
				fmt.Println(animal.food)
			case "move":
				animal.Move()
				fmt.Println(animal.locomotion)
			case "speak":
				animal.Speak()
				fmt.Println(animal.noise)
			default:
				fmt.Println("action not supported yet")
			}
		} else {
			fmt.Println("input error")
		}
	}
}
