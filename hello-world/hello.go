package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	fmt.Println(Hello("bjarki", "ru"))
}

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
	//etc..
}

func greeting(language string) string {
	greeting, exists := greetings[language]

	if exists {
		return greeting
	}

	return "Hello"
}