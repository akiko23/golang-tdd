package main

import (
  "testing"
  "sync"
)

func NewCounter() *Counter {
	return &Counter{}
}


func TestCounter(t *testing.T) {
  t.Run("Test counter in single thread", func(t *testing.T) {
    counter := NewCounter()
    counter.Inc()
    counter.Inc()
    counter.Inc()
       
    assertCounter(t, counter, 3)    
  })

  t.Run("Counter in multiple threads", func(t *testing.T) {
    wantedCount := 1000
    counter := &Counter{}
  
    var wg sync.WaitGroup
    wg.Add(wantedCount)

    for i := 0; i < wantedCount; i++ {
      go func() {
        counter.Inc()
        wg.Done()
      }()
    }
    wg.Wait()

    assertCounter(t, counter, wantedCount)
  })
}

func assertCounter(t *testing.T, c *Counter, want int) {
  t.Helper()

  got := c.Value()
  if got != want {
    t.Errorf("Want %d but got %d", want, got)
  }
}
