package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	o := 0.0
	for z != o {

		o = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z

}

func main() {
	fmt.Println(Sqrt(9))
}
