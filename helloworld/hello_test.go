package helloworld

import "testing"

func TestHello(t *testing.T) {
  t.Run("Say hello to people", func(t *testing.T) {
    got := Hello("John", "")
    expected := "Hello, John"
    assertCorrectMessage(t, got, expected)
  })

  t.Run("Say hello to world", func(t *testing.T) {
    got := Hello("", "")
    expected := "Hello, World"
    assertCorrectMessage(t, got, expected)
  })

  t.Run("Greetings in Spanish", func(t *testing.T) {
    got := Hello("Ricardo", "sp")
    expected := "Hola, Ricardo"
    assertCorrectMessage(t, got, expected)
  })

  t.Run("Greetings in French", func(t *testing.T) {
    got := Hello("Lui", "fr")
    expected := "Bonjour, Lui"
    assertCorrectMessage(t, got, expected)
  })
}


func assertCorrectMessage(t testing.TB, got, expected string) {
  t.Helper()
  if got != expected {
    t.Errorf("Got %q expected %q", got, expected)
  }
}


