package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f",
		e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), ErrNegativeSqrt(x) // Coul also return 0 here, if you want to go without the math library
	}
	
	z := 1.0
	z -= (z*z - x) / (2*z)
	next := z - (z*z - x) / (2*z)
	
	for i := 0; z > next; i++ {
		z -= (z*z - x) / (2*z)
		next -= (z*z - x) / (2*z)
	}
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}