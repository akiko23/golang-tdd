package racer

import (
  "time"
  "errors"
  "net/http"
)

var (
  racerTimeout = 10 * time.Second
  ErrTimeOut = errors.New("time out waiting for the servers")
)

func Racer(a, b string) (winner string, err error) {
  select {
  case <-ping(a):
    return a, nil
  case <-ping(b):
    return b, nil
  case <-time.After(10 * time.Second):
    return "", ErrTimeOut 
  }
}


func ping(url string) chan struct{} {
  ch := make(chan struct{})
  go func() {
    http.Get(url)
    close(ch)
  }()
  return ch
}

func measureResponseTime(url string) time.Duration {
  start := time.Now()
  http.Get(url)
  duration := time.Since(start)

  return duration
}

