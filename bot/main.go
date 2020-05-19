package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreating(b bot) {
	fmt.Printf("%s\n", b.getGreeting())
}

type spanishBot struct{}
type englishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)
}
