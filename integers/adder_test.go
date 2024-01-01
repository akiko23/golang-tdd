package integers

import (
  "testing"
  "fmt"
)


func TestAdder(t *testing.T) {
  //t.Run("Add positive numbers", func(t *testing.T) {
    //got := Add(4, 5)
    //expected := 9
    //assertSum(t, got, expected)
  //})

  got := Add(4, 5)
  expected := 9

  if got != expected {
    t.Errorf("Expected %d but got %d", expected, got)
  }
}

func ExampleAdd() {
  res := Add(2, 4)
  fmt.Println(res)
  // Output: 6
}

