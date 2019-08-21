package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct {
	message string
}

type spanishBot struct {
	mensaje string
}

func main() {

	//	fmt.Println(newEnglishBot("Hello world!").getGreeting())
	//	fmt.Println(newSpanishBot("¡Hola mundo!").getGreeting())

	eBot := newEnglishBot("Hello, how are you doing, man?")
	sBot := newSpanishBot("Hola, cómo te va, tío?")

	printGreeting(eBot)

	printGreeting(sBot)

}

func newEnglishBot(msg string) englishBot {
	return (englishBot{message: msg})
}

func newSpanishBot(mnsj string) spanishBot {
	return (spanishBot{mensaje: mnsj})

}

func (eb englishBot) getGreeting() string {
	//specific logic to thistype

	return eb.message

}

func (sb spanishBot) getGreeting() string {
	//specific logic to thistype

	return sb.mensaje

}

func printGreeting(ib bot) {
	fmt.Println(ib.getGreeting())
}
