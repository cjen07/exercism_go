package accumulate

const testVersion = 1

func Accumulate(d []string, f func(string) string) []string {
  ss := make([]string, len(d))
  for i, v := range d {
    ss[i] = f(v)
  }
  return ss
}