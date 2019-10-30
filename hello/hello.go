package main

import "fmt"

const englishHelloPrefix = "Hello "

// Hello function returns a greeting when called
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("David"))
}
