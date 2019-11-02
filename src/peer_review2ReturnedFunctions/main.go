package main

import (
	"fmt"
	"math"
	"strconv"
)

func AskForVariable(prompt string) (float64, error) {
	var userInput string
	fmt.Println(prompt)
	fmt.Scan(&userInput)

	askedVariable, err := strconv.ParseFloat(userInput, 64)

	return askedVariable, err
}

func GenDisplaceFn(zeroVelocity, zeroDisplacement, acceleration float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return ((0.5 * acceleration) * math.Pow(t, 2)) + (zeroVelocity * t) + zeroDisplacement
	}
	return fn
}

func main() {
	acceleration, errAcceleration := AskForVariable("Please enter an acceleration float64")
	initialVelocity, errVelocity := AskForVariable("Please enter an initial velocity float64")
	initialDisplacement, errDisplacement := AskForVariable("Please enter an initial displacement float64")

	if errAcceleration == nil && errVelocity == nil && errDisplacement == nil {
		time, errTime := AskForVariable("Please enter time")
		if errTime == nil {
			fn := GenDisplaceFn(initialVelocity, initialDisplacement, acceleration)
			fmt.Println("Displacement: ")
			fmt.Println(fn(time))
		} else {
			fmt.Println("Please check your input")
		}
	} else {
		fmt.Println("Please check your input")
	}
}
