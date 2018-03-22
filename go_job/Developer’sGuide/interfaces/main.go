package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}

type bot interface {
	getGreeting() string
}

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot)getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot)getGreeting() string {
	return "Hola!"
}