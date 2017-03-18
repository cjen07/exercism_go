package triangle

import "math"

const testVersion = 3

// Notice it returns this type.  Pick something suitable.
type Kind int

// Pick values for the following identifiers used by the test program.
const NaT Kind = 0 // not a triangle
const Equ Kind = 1 // equilateral
const Iso Kind = 2 // isosceles
const Sca Kind = 3 // scalene

// Code this function.
func KindFromSides(a, b, c float64) Kind {
  if IsNum(a) && IsNum(b) && IsNum(c) {
    if a <= 0 || b <= 0 || c <= 0 {
      return NaT
    } else if a == b && b == c {
      return Equ
    } else if (a + b < c) || (a + c < b) || (b + c < a) {
      return NaT
    } else if a == b || b == c || a == c {
      return Iso
    } else {
      return Sca
    }
  } else {
    return NaT
  }
}

func IsNum(a float64) bool {
  return !(math.IsInf(a, 0) || math.IsNaN(a))
}