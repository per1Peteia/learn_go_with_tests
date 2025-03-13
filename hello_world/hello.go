package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const spanishPrefix = "Hola"
const englishPrefix = "Hello"
const frenchPrefix = "Bonjour"

func Hello(name, language string) string {
	if name == "" {
		return "Hello, World!"
	}
	if language == spanish {
		return fmt.Sprintf("%s, %s!", spanishPrefix, name)
	}
	if language == french {
		return fmt.Sprintf("%s, %s!", frenchPrefix, name)
	}
	return fmt.Sprintf("%s, %s!", englishPrefix, name)
}

func main() {
	fmt.Println(Hello("Justus", ""))
}
