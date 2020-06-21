package main

import "fmt"

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour, "
const italianGreeting = "Buongiorno, "
const spanishGreeting = "Hola, "
const french = "French"
const italian = "Italian"
const spanish = "Spanish"

// Hello returns a greeting along with a provided name
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greeting(language) + name
}

func greeting(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	case italian:
		greeting = italianGreeting
	default:
		greeting = englishGreeting
	}
	return
}

func main() {
	fmt.Println("hello")
}
