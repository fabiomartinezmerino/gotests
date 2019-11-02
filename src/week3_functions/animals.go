package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//Define Animal Struct
type Animal struct {
	name, food_eaten, locomotion_method, spoken_sound string
}

//Define methods associated with Animal

func (a *Animal) Eat() {

	fmt.Printf("%v eats %v\n", a.name, a.food_eaten)

}

func (a *Animal) Move() {

	fmt.Printf("%v locomotion method is: %v\n", a.name, a.locomotion_method)

}

func (a *Animal) Speak() {

	fmt.Printf("%v sound is: %v\n", a.name, a.spoken_sound)

}

func main() {

	//Instantiate three animals
	cow := Animal{name: "cow", food_eaten: "grass", locomotion_method: "walk", spoken_sound: "moo"}
	bird := Animal{name: "bird", food_eaten: "worms", locomotion_method: "fly", spoken_sound: "peep"}
	snake := Animal{name: "snake", food_eaten: "mice", locomotion_method: "slither", spoken_sound: "hsss"}

	//Allow the user to enter animal and ask question and answer it - infinite loop you need to kill process to stop execution
	var animal, question string

	for {
		//Prompt
		fmt.Println("Please, introduce animal & ask question > ")

		reader := bufio.NewReader(os.Stdin)
		//Read from console and store into string
		str, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		strSl := strings.Fields(str)

		animal, question = strings.ToLower(strSl[0]), strings.ToLower(strSl[1])

		switch animal {
		case "cow":
			switch question {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Println("Please ask again")
			}

		case "bird":
			switch question {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			default:
				fmt.Println("Please ask again")
			}
		case "snake":
			switch question {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Println("Please ask again")
			}

		default:
			fmt.Println("Please ask again")

		}
	}

}
