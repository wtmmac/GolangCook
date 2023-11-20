package hello

import "fmt"

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
	// etc..
}

func greeting(language string) string {
	greeting, exists := greetings[language]
	if exists {
		return greeting
	}
	return "Hello"
}
