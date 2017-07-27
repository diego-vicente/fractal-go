package main

import (
	"fmt"
	"image"
)

// Define the type of a FractalGenerator
type FractalGenerator func(int) *image.RGBA

// MandelbrotNaive creates a Mandelbrot fractal sequentially
func MandelbrotNaive() (fractal *image.RGBA) {
	step := ComputeStep()
	ComputeYBounds(step)
	fractal = image.NewRGBA(image.Rect(0, 0, XSIZE, YSIZE))

	for i := 0; i <= YSIZE; i++ {
		for j := 0; j < XSIZE; j++ {
			n := ComplexAt(i, j, step)
			iterations := ComputeIterations(n)
			fractal.Set(j, i, FancyColor(iterations))
		}
	}

	return fractal
}

// MandelbrotBands creates a Mandelbrot fractal launching a number of
// routines, each one in charge of a horizontal slice of the image.
func MandelbrotBands(routines int) (fractal *image.RGBA) {
	step := ComputeStep()
	ComputeYBounds(step)
	fractal = image.NewRGBA(image.Rect(0, 0, XSIZE, YSIZE))
	bandSize := YSIZE / routines
	done := make(chan int)

	for i := 0; i < routines; i++ {
		go ComputeBand(i*bandSize, (i+1)*bandSize, step, fractal, done)
	}

	for i := 0; i < routines; i++ {
		<-done
	}

	return fractal
}

// ComputeBand computes a horizontal slice of the Mandelbrot image that is
// bounded by initialY and finalY parameters. It receives a channel that is
// used to poll when the job is done.
func ComputeBand(initialY, finalY int, step float64, fractal *image.RGBA,
	done chan int) {

	for i := initialY; i <= finalY; i++ {
		for j := 0; j < XSIZE; j++ {
			n := ComplexAt(i, j, step)
			iterations := ComputeIterations(n)
			fractal.Set(j, i, FancyColor(iterations))
		}
	}

	done <- 1
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
