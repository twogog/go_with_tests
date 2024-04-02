package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
	frenchHelloPrefix  = "Bonjour, "
	russian            = "Russian"
	french             = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case russian:
		prefix = russianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
