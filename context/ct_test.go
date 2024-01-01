package context

import (
  "testing"
  "net/http"
  "net/http/httptest"
)

type StubStore struct {
  response string
}


func (s *StubStore) Fetch() string {
  return s.response
}

func (s *StubStore) Cancel() {}


func TestServer(t *testing.T) {
  data := "Hello, World!"
  store := &SpyStore{response: data}

  svr := Server(store)

  request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}


