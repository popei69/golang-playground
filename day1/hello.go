package main

import "fmt"

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour "

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}

	prefix := MakeGreeting(language)
	return prefix + name
}

func MakeGreeting(language string) string {
	switch language {
	case "French":
		return frenchGreeting
	default:
		return englishGreeting
	}
}

func main() {
	fmt.Println(Hello("", "English"))
}
