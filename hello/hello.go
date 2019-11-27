package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const spanish = "Spanish"
const french = "French"
const world = "World"

func main() {
	fmt.Println(Hello("world", "English"))
}

func Hello(name, lang string) string {
	if name == "" {
		name = world
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
