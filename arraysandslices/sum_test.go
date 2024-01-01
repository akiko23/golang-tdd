package arraysandslices

import (
  "testing"
  "reflect"
)

func TestSum(t *testing.T) {
  
  
  assertCollectionSum := func (t testing.TB, got, expected int) {
    t.Helper()
    if got != expected {
      t.Errorf("Expected %d but got %v", expected, got)
    }
  }

  t.Run("Collection of size 5", func(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    expected := 15

    assertCollectionSum(t, got, expected) 
  })

  t.Run("Collection of any size", func(t *testing.T) {
    numbers := []int{1, 2, 3, 4}
    
    got := Sum(numbers)
    expected := 10

    assertCollectionSum(t, got, expected)
  })
}


func TestSumAll(t *testing.T) {
  
  assertSumAll := func(t testing.TB, got, expected []int) {
    t.Helper()
    if !reflect.DeepEqual(got, expected) {
      t.Errorf("Expected %v but got %v", expected, got)
    }
  }

  t.Run("Non empty slices", func(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{1, 4, 7})
    expected := []int{3, 12}
    assertSumAll(t, got, expected)
    })

  t.Run("An empty slice", func(t *testing.T) {
    got := SumAll([]int{}, []int{1, 2, 6})
    expected := []int{0, 9}
    assertSumAll(t, got, expected)
  })
}

