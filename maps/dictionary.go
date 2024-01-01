package main

import "errors"

var (
  ErrWordNotFound = errors.New("could not find the word you are looking for")
  ErrWordExists = errors.New("word you are trying to add already exists")
  ErrWordDoesNotExist = errors.New("word you're trying update does not exist")
  ErrWordForDeleteDoesNotExist = errors.New("word you're trying delete doesn't exist")
)

func (d Dictionary) Search(word string) (string, error) {
  definition, ok := d[word]
  if !ok {
    return "", ErrWordNotFound
  }

  return definition, nil
}


func (d Dictionary) Add(word, definition string) error {
  _, err := d.Search(word)
  
  switch err {
  case ErrWordNotFound:
    d[word] = definition
  case nil:
    return ErrWordExists
  default:
    return err
  }

  return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
  _, err := d.Search(word)
  if err == ErrWordNotFound {
    return ErrWordDoesNotExist
  }

  d[word] = newDefinition
  return nil  
}


func (d Dictionary) Delete(word string) error {
  _, err := d.Search(word)
  if err == ErrWordNotFound {
    return ErrWordForDeleteDoesNotExist
  }

  delete(d, word)
  return nil
}

