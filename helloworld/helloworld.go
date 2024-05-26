package helloworld

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

type Lang string

const (
	LangEnglish Lang = "English"
	LangSpanish Lang = "Spanish"
	LangFrench  Lang = "French"
)

func Hello(name string, language Lang) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case LangSpanish:
		prefix = spanishHelloPrefix
	case LangFrench:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}
