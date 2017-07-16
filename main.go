package main

import (
	"fmt"
	"math/cmplx"
)

const MAX_ITER = 254

func main() {
	z := 0.1 + 0.1i
	fmt.Println(ComputeIterations(z))
}

// ComputeIterations returns how many iterations took the complex n to diverge
func ComputeIterations(n complex128) (iterations uint8) {
	iterations = 0
	z := n

	for cmplx.Abs(z) < 4 && iterations <= MAX_ITER {
		z = cmplx.Pow(z, 2) + n
		iterations += 1
	}

	return iterations
}
