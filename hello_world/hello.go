package main

import (
	"fmt"
)

func Hello(str string) string {
	if str == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", str)
}

func main() {
	fmt.Println(Hello("Justus"))
}
