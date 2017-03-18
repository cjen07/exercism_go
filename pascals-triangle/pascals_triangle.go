package pascal

const testVersion = 1

func MapReduce(a []int, init int, f func(int, int) (int, int)) []int {
  result := make([]int, len(a))
  acc := init
  for i, v := range a {
    result[i], acc = f(v, acc)
  }
  return result
}

func DoTriangle(a []int) []int {
  return MapReduce(a, 0, func(v, acc int) (int, int) { return v + acc, v })
}

func Triangle(d int) [][]int {
  aa := make([][]int, d)
  i := 0
  a := make([]int, 0)
  for i < d {
    a = append(DoTriangle(a), 1)
    aa[i] = a
    i++
  }
  return aa
}