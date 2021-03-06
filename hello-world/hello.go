package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	fmt.Println(Hello("bjarki", "ru"))
}

const englishHelloPrefix = "Hello"

func Hello(name, language string) string {
	return fmt.Sprintf(
		"%s, %s",
		greeting(language),
		name,
	)
}

var greetings = map[string]string{
	"es": "Hola",
	"fr": "Bonjour",
	"ru": "Привет",
	"ja": "こんにちは",
	"de": "Guten Tag",
	//etc..
}

func greeting(language string) string {
	greeting, exists := greetings[language]

	if exists {
		return greeting
	}

	return englishHelloPrefix
}
