package main

import "fmt"

const englishGreeting = "Hello, "

func Hello(name string) string {
	if len(name) > 0 {
		return englishGreeting + name
	} else {
		return "Hello, World"
	}

}

func main() {
	fmt.Println(Hello(""))
}
