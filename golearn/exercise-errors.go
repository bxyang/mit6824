

package main

import (
	"fmt"
	"math"
)

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
	} else {
		return "OK"
	}
}

func Sqrt(x float64) (float64, error) {
	return math.Sqrt(x), ErrorNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
