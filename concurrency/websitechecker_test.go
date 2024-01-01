package concurrency

import (
  "testing"
  "time"
  "reflect"
)

const IncorrectURL = "wafoa:\\falsglaglka.com"


func MockWebsiteChecker(url string) bool {
  return url != IncorrectURL 
}


func TestCheckWebsites(t *testing.T) {
  websites := []string{
    IncorrectURL,
    "https://google.com",
    "https://yandex.ru",
  }

  want := map[string]bool{
    IncorrectURL:         false,
    "https://google.com": true,
    "https://yandex.ru":  true,

  }

  got := CheckWebsites(MockWebsiteChecker, websites)

  if !reflect.DeepEqual(got, want) {
    t.Fatalf("Want %v but got %v", want, got)
  }
}


func slowStubWebsiteChecker(_ string) bool {
  time.Sleep(20 * time.Millisecond)
  return true
}

func BenchmarkCheckWebsites(b *testing.B) {
  urls := make([]string, 100)
  for i := 0; i < len(urls); i++ {
    urls[i] = "a url"
  }
  b.ResetTimer()
  
  for i := 0; i < b.N; i++ {
    CheckWebsites(slowStubWebsiteChecker, urls)
  }
}



