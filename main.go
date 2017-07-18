package main

import (
	"fmt"
	"math/cmplx"
)

// Constants for the image size.
const XSIZE = 2560
const YSIZE = 2048

// Constant for the number of iterations to perform.
const MAX_ITER = 255

// Parameters for generating the image.
const xLeft = -2.0
const xRight = 1.0
const yCenter = 0.0

func main() {
	z := 0.1 + 0.1i
	fmt.Println(ComputeIterations(z))
}

// ComputeIterations returns how many iterations took the complex n to diverge.
func ComputeIterations(n complex128) (iterations uint8) {
	z := n
	for cmplx.Abs(z) < 4 && iterations < MAX_ITER {
		z = cmplx.Pow(z, 2) + n
		iterations += 1
	}
	return iterations
}

// ComputeStep returns the imaginary step from a pixel to another.
func ComputeStep() float64 {
	return (xRight - xLeft) / XSIZE
}

// ComputeYBounds returns the y-axis boundaries for a given step.
func ComputeYBounds(step float64) (yUpper, yLower float64) {
	yUpper = yCenter + (step*YSIZE)/2
	yLower = yCenter - (step*YSIZE)/2
	return yUpper, yLower
}
