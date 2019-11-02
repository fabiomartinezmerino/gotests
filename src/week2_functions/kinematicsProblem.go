package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//Function to parse the string and extract float64 values
func parseFloats(str string) ([]float64, error) {
	var (
		fields           = strings.Fields(str)
		float_parameters = make([]float64, len(fields))
		err              error
	)
	for i, v := range fields {
		float_parameters[i], err = strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
	}
	return float_parameters, nil
}

func GenDisplaceFn(a, v, si float64) func(float64) float64 {

	//lambda function that takes in a,v and si from enclosing function GenDisplaceFn and returns space covered as a function of time.
	//lambda is stored in f that is returned afterwards
	f := func(t float64) float64 {
		return (0.5*a*math.Pow(t, 2) + v*t + si)

	}

	return f

}

func main() {

	//Prompt for a, v0, s0 and t
	fmt.Println("Please, introduce the following parameters separate by spaces: acceleration, initial velocity, initial displacement, time")

	reader := bufio.NewReader(os.Stdin)
	//Read from console and store into string
	strParameters, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	//Parse string to obtain a slice containing parameters as float64 values - this function is custom built
	slParameters, err := parseFloats(strParameters)

	if err != nil {
		log.Fatal(err)
	}

	//for debugging purposes: makes sure all values are properly loaded
	/*
		for p, v := range slParameters {

			fmt.Printf("Index: %v - Value: %v\n", p, v)
		}
	*/

	//we declare a function to store the lambda function that computes space as function of time
	//we can find acceleration in slParameters[0], initial velocity in slParameters[1], initial displacement in slParameters[2]
	space := GenDisplaceFn(slParameters[0], slParameters[1], slParameters[2])
	//and now we call the returned function paasing time in and print the result of the kinematicks problem for input initial conditions and parameters
	totalDisplacement := space(slParameters[3])

	fmt.Println("--- Initial conditions and parameters ---")
	fmt.Printf("Acceleration: %v m/s2\n", slParameters[0])
	fmt.Printf("Initial velocity %v m/s\n", slParameters[1])
	fmt.Printf("initial displacement %v m\n", slParameters[2])
	fmt.Printf("Elapsed time you want to compute total displacement for %v s\n", slParameters[3])
	fmt.Printf("Total displacement %v m\n", totalDisplacement)
}
