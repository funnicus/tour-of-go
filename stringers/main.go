package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	z -= (z*z - x) / (2*z)
	next := z - (z*z - x) / (2*z)
	
	for i := 0; z > next; i++ {
		z -= (z*z - x) / (2*z)
		next -= (z*z - x) / (2*z)
	}
	
	return
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}