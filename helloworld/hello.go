package helloworld

import "fmt"


const (
  englishLanguageCode = "en"
  spanishLanguageCode = "sp"
  frenchLanguageCode = "fr"

  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
  frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language_code string) string {
  if name == "" {
    name = "World"
  } 
  
	return greetingPrefix(language_code) + name
}


func greetingPrefix(language_code string) (helloPrefix string) { 
  switch language_code {
  case frenchLanguageCode:
    helloPrefix = frenchHelloPrefix 
  case spanishLanguageCode:
    helloPrefix = spanishHelloPrefix
  default:
    helloPrefix = englishHelloPrefix
  }

  return helloPrefix
}

func main() {
	fmt.Println(Hello("Geyb", ""))
}

