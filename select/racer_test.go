package racer

import (
  "testing"
  "time"
  "net/http/httptest"
  "net/http"
)


func TestRacer(t *testing.T) {
  
  t.Run("check the fastest server", func(t *testing.T) {
    slowServer := makeDelayedServer(20 * time.Millisecond)
	  fastServer := makeDelayedServer(0 * time.Millisecond)	
  
    defer slowServer.Close()
    defer fastServer.Close()

	  slowURL := slowServer.URL
	  fastURL := fastServer.URL

	  want := fastURL
	  got, _ := Racer(slowURL, fastURL)

	  if got != want {
		  t.Errorf("got %q, want %q", got, want)
	  }
  })

  t.Run("check an error on responding for more then 10s", func(t *testing.T) {
    serverA := makeDelayedServer(11 * time.Second)
    serverB := makeDelayedServer(12 * time.Second)

    defer serverA.Close()
    defer serverB.Close()

    _, err := Racer(serverA.URL, serverB.URL)
    if err == nil {
      t.Error("Expected an error")
    }
  }) 
}


func makeDelayedServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}


