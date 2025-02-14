package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"

	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	enlishHelloPrefix  = "Hello, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(langauge string) (prefix string) {
	switch langauge {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = enlishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
