package reflection

import (
  "testing"
)

func TestWalk(t *testing.T) {
  t.Run("walk for the struct with on string field", func(t *testing.T) {
    expected := "Chris"
    got := []string{}

    myStruct := struct {
      name string
    }{expected}
    
    walk(myStruct, func(input string) {
      got = append(got, input)
    })
    
    if len(got) != 1 {
		  t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	  }

    if got[0] != expected {
      t.Errorf("want '%s' but got '%s'", expected, got[0])
    }
  })

}


