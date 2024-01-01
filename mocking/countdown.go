package main

import( 
  "fmt"
  "io"
  "time"
)

type Sleeper interface {
  Sleep()
}

type SpySleeper struct {
  Calls int
}

func (s *SpySleeper) Sleep() {
  s.Calls++
}



func Countdown(writer io.Writer, spySleeper Sleeper) {
  for i := 3; i > 0; i-- {
    fmt.Fprintf(writer, "%d\n", i)
    time.Sleep(1 * time.Second)

    spySleeper.Sleep()
  }
  fmt.Fprint(writer, "Go!")
}


