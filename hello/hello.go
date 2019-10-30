package main

import "fmt"

const englishPrefix = "Hello "
const spanishPrefix = "Ola! "
const frenchPrefix = "Bonjour "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchPrefix
	case "spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

// Hello function returns a greeting when called
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("David", ""))
}
