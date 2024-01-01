package main

import (
  "bytes"
  "testing"
)

func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}
  spySleeper := &SpySleeper{}

  Countdown(buffer, spySleeper)
  
  got := buffer.String()
  want := 
  `3
2
1
Go!`
  

  expectedCalls := 3 

  if got != want {
    t.Errorf("want '%s' but got '%s'", want, got)
  }

  if spySleeper.Calls != expectedCalls {
    t.Errorf("want %d calls but got %d calls", expectedCalls, spySleeper.Calls) 
  }


}

