package iteration

func Repeat(st string, repeatCount int) (repeated string) {
  for i := 0; i < repeatCount; i++ {
    repeated += st
  }
  return repeated
}

