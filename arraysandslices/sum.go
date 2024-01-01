package arraysandslices


func Sum(numbers []int) (sum int) {
  for _, number := range numbers {
    sum += number
  }
  return sum
}


func SumAll(numbersToSum ...[]int) []int { 
  totals := []int{}

  for _, numbers := range numbersToSum {
    totals = append(totals, Sum(numbers))
  }
  return totals
}

