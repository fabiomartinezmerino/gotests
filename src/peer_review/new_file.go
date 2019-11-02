package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	first string
	last  string
}

func main() {
	fmt.Println("Enter the name of the text file:")
	consoleReader := bufio.NewReader(os.Stdin)
	fileName, err := consoleReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	names := make([]name, 0, 3)

	fileName = strings.Trim(fileName, "\n")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var naam name
		naam.first, naam.last = s[0], s[1]
		names = append(names, naam)

	}

	for _, v := range names {
		fmt.Println(v)
	}
}
