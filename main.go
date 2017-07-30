package main

import (
	//	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"time"
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

// The y axis parameter have to be computed depending on the size of the image.
var yUpper, yLower float64

// Flag to obtain the number of goroutines to use in the generation
var NRoutines = flag.Int("n", 1, "number of goroutines to launch")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*NRoutines)

	start := time.Now()
	fractal := MandelbrotLines(*NRoutines)
	end := time.Now()

	fmt.Printf("Generation took: %s\n", end.Sub(start))
	SaveImage(fractal)
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

// SaveImage saves the created fractal representations
func SaveImage(fractal *image.RGBA) {
	f, _ := os.Create("mandelbrot.png")
	defer f.Close()
	if err := png.Encode(f, fractal); err != nil {
		panic(err)
	}
}

func FancyColour(iter uint8) color.RGBA {
	switch {
	case iter == MAX_ITER:
		return color.RGBA{0, 0, 0, 255}
	case iter < 8:
		return color.RGBA{128 + (iter * 16), 0, 0, 255}
	case iter < 24:
		return color.RGBA{255, (iter - 8) * 16, (iter - 8) * 16, 255}
	case iter < 160:
		return color.RGBA{255 - (iter-24)*2, 255 - (iter-24)*2, 255, 255}
	default:
		return color.RGBA{(iter - 160) * 2, (iter - 160) * 2,
			255 - (iter-160)*2, 255}
	}
}

// SimpleColor is a function used in NTNU to draw a BMP image of the Mandelbrot
// fractal.
func SimpleColor(iter uint8) color.RGBA {
	switch {
	case iter == MAX_ITER:
		return color.RGBA{0, 0, 0, 255}
	case iter < 8:
		return color.RGBA{128 + (iter * 16), 0, 0, 255}
	case iter < 24:
		return color.RGBA{255, (iter - 8) * 16, (iter - 8) * 16, 255}
	case iter < 160:
		return color.RGBA{255 - (iter-24)*2, 255 - (iter-24)*2, 255, 255}
	default:
		return color.RGBA{(iter - 160) * 2, (iter - 160) * 2,
			255 - (iter-160)*2, 255}
	}
}

// FancyColor tries to mimic the interpolation used to generate the Mandelbrot
// image in Wikipedia linearly. As found in StackOverflow, this color palette
// is defined by these control points:
// Position = 0.0     Color = (0,   7,   100)
// Position = 0.16    Color = (32,  107, 203)
// Position = 0.42    Color = (237, 255, 255)
// Position = 0.6425  Color = (255, 170, 0)
// Position = 0.8575  Color = (0,   2,   0)
func FancyColor(iter uint8) color.RGBA {
	position := float64(iter) / float64(MAX_ITER)
	switch {
	case position == 1.0:
		return color.RGBA{0, 0, 0, 255}
	case position < 0.16:
		offset := position / 0.16
		return color.RGBA{
			uint8(0 + 32*offset),
			uint8(7 + 100*offset),
			uint8(100 + 103*offset),
			255}
	case position < 0.42:
		offset := (position - 0.16) / 0.26
		return color.RGBA{
			uint8(32 + 205*offset),
			uint8(107 + 148*offset),
			uint8(203 + 52*offset),
			255}
	case position < 0.6425:
		offset := (position - 0.42) / 0.2225
		return color.RGBA{
			uint8(237 + 18*offset),
			uint8(255 - 85*offset),
			uint8(255 - 255*offset),
			255}
	case position < 0.8575:
		offset := (position - 0.6425) / 0.2150
		return color.RGBA{
			uint8(255 - 255*offset),
			uint8(170 - 168*offset),
			0, 255}
	default:
		offset := (position - 0.8575) / 0.1425
		return color.RGBA{0, 0, uint8(0 + 2*offset), 255}
	}
}
