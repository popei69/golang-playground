package main

import "fmt"

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour "

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}

	return MakeGreeting(language) + name
}

func MakeGreeting(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("", "English"))
}
