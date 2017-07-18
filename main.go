package main

import (
	"fmt"
	"math/cmplx"
)

// Constants for the image size.
const XSIZE = 50 //2560
const YSIZE = 40 // 2048

// Constant for the number of iterations to perform.
const MAX_ITER = 255

// Parameters for generating the image.
const xLeft = -2.0
const xRight = 1.0
const yCenter = 0.0

// The y axis parameter have to be computed depending on the size of the image.
var yUpper, yLower float64

func main() {
	PrintFractal()
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

// ComputeYBounds sets the global y-axis boundaries for a given step.
func ComputeYBounds(step float64) {
	yUpper = yCenter + (step*YSIZE)/2
	yLower = yCenter - (step*YSIZE)/2
}

// ComplexAt returns the associate complex to a i-j iteration and a step.
func ComplexAt(i, j int, step float64) (n complex128) {
	return complex(xLeft+float64(j)*step, yUpper-float64(i)*step)
}

// PrintFractal displays the fractal on the screen, using text
func PrintFractal() {
	step := ComputeStep()
	ComputeYBounds(step)

	for i := 0; i < 40; i++ {
		for j := 0; j < 50; j++ {
			n := ComplexAt(i, j, step)
			iterations := ComputeIterations(n)
			if iterations > MAX_ITER/10 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
