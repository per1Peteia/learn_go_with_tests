package main

import (
	"fmt"
)

const (
	spanish       = "Spanish"
	french        = "French"
	german        = "German"
	spanishPrefix = "Hola"
	englishPrefix = "Hello"
	frenchPrefix  = "Bonjour"
	germanPrefix  = "Hallo"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case german:
		prefix = germanPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Justus", ""))
}
