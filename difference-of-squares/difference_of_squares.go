package diffsquares

const testVersion = 1

func Sum(a []int) int {
  result := 0
  for _, v := range a {
    result += v
  }
  return result
}

func Square(d int) int {
  return d * d
}

func Map(a []int, f func(int) int) []int {
  result := make([]int, len(a))
  for i, v := range a {
    result[i] = f(v)
  }
  return result
}

func MakeArray(d int) []int {
  result := make([]int, d)
  i := 0
  for i < d {
    result[i] = i + 1
    i++
  }
  return result
}

func SquareOfSums(d int) int {
  a := MakeArray(d)
  return Square(Sum(a))
}

func SumOfSquares(d int) int {
  a := MakeArray(d)
  return Sum(Map(a, Square))
}

func Difference(d int) int {
  return SquareOfSums(d) - SumOfSquares(d)
}