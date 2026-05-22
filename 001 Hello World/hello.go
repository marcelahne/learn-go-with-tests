package main

import "fmt"

const (
	german  = "German"
	french  = "French"
	spanish = "Spanish"

	defaultPrefix = "Hello, "
	germanPrefix  = "Hallo, "
	frenchPrefix  = "Bonjour, "
	spanishPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case german:
		return germanPrefix
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	}

	return defaultPrefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
