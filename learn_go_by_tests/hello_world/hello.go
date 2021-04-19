package main

import "fmt"

const (
	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	spanishLang          = "ES"
	frenchHelloPrefix    = "Bonjour, "
	frenchLang           = "FR"
	slovenianHelloPrefix = "Zdravo, "
	slovenianLang        = "SI"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanishLang:
		return spanishHelloPrefix
	case frenchLang:
		return frenchHelloPrefix
	case slovenianLang:
		return slovenianHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Domen", ""))
}
