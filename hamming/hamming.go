package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	// your code here!
	la := len(a)
	lb := len(b)

	if la != lb {
		return -1, errors.New("not match")
	} else {
		i := 0
		result := 0
		for i < la {
			if a[i] != b[i] {
				result++
			}
			i++
		}
		return result, nil
	}
}
