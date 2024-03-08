package main

const (
	Spanish = "Spanish"
	French  = "French"
	Basque  = "Basque"

	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	basqueHelloPrefix  = "Kaixo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Basque":
		prefix = basqueHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
