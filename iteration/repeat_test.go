package iteration

import (
  "testing"
  "fmt"
)


func TestRepeat(t *testing.T) {
  t.Run("5 repeats", func(t *testing.T) {
    got := Repeat("a", 5)
    expected := "aaaaa"
    assertRepeatedString(t, got, expected)  
  })
  
  t.Run("Zero repeats", func(t *testing.T) {
    got := Repeat("a", 0)
    expected := ""
    assertRepeatedString(t, got, expected)
  })
}

func assertRepeatedString(t testing.TB, got, expected string) {
  t.Helper()
  if got != expected {
   t.Errorf("Expected %q but got %q", expected, got)
  }
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
  repeated := Repeat("hello", 2)
  fmt.Println(repeated)
  // Output: hellohello
}


