package main

import "testing"



func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}
  
  t.Run("word is in dict", func(t *testing.T) {
    got, _ := dictionary.Search("test")
    want := "this is just a test"
    
    assertStrings(t, got, want)
  })

  t.Run("word is not in dict", func(t *testing.T) {
    _, err := dictionary.Search("unknown") 
    if err == nil {
      t.Fatal("Expected to get an error")
    }

    assertStrings(t, err.Error(), ErrWordNotFound.Error())
  })  
}


func TestAdd(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("new word", func(t *testing.T) {
    word := "test2"
    definition := "there is a new word in dict"
    
    dictionary.Add(word, definition)
   
    assertDefinition(t, dictionary, word, definition) 
  })

  t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "new definition for test"

    err := dictionary.Add(word, definition)
    assertError(t, err, ErrWordExists)
  })
}

func TestUpdate(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("update definition of existing word", func(t *testing.T) {
    word := "test"
    newDefinition := "this is not just a test"

    err := dictionary.Update(word, newDefinition)
  
    assertError(t, err, nil)
    assertDefinition(t, dictionary, word, newDefinition)
  })
  
  t.Run("update definition of non existing word", func(t *testing.T) {
    word := "test"
    definition := "this is not just a test"

    err := dictionary.Update(word, definition)
    
    assertError(t, err, nil)
  }) 
}

func TestDelete(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("delete existing word", func(t *testing.T) {
    word := "test"

    err := dictionary.Delete(word) 
    assertError(t, err, nil)
  })

  t.Run("delete non existing word", func(t *testing.T) {
    word := "test2"

    err := dictionary.Delete(word)
    assertError(t, err, ErrWordForDeleteDoesNotExist)
  })
}



func assertStrings(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("Want '%s' but got '%s'", want, got)
  }
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    got, err := dictionary.Search(word) 
    
    if err != nil {
      t.Fatalf("should find added word: %s", err)
    }
    assertStrings(t, got, definition)
}


