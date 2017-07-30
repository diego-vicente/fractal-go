package main

import (
	"os"
	"testing"
)

// Map of imaginary numbers and its iterations
var iterationTests = map[complex128]uint8{
	(1 - 2i):      1,
	(1 - 1.5i):    1,
	(1 - 1i):      2,
	(1 - 0.5i):    2,
	(1 + 0i):      2,
	(1 + 0.5i):    2,
	(1 + 1i):      2,
	(1 + 1.5i):    1,
	(1 + 2i):      1,
	(0.5 - 2i):    1,
	(0.5 - 1.5i):  2,
	(0.5 - 1i):    3,
	(0.5 - 0.5i):  5,
	(0.5 + 0i):    5,
	(0.5 + 0.5i):  5,
	(0.5 + 1i):    3,
	(0.5 + 1.5i):  2,
	(0.5 + 2i):    1,
	(0 - 2i):      1,
	(0 - 1.5i):    2,
	(0 - 1i):      255,
	(0 - 0.5i):    255,
	(0 + 0i):      255,
	(0 + 0.5i):    255,
	(0 + 1i):      255,
	(0 + 1.5i):    2,
	(0 + 2i):      1,
	(-0.5 - 2i):   1,
	(-0.5 - 1.5i): 2,
	(-0.5 - 1i):   4,
	(-0.5 - 0.5i): 255,
	(-0.5 + 0i):   255,
	(-0.5 + 0.5i): 255,
	(-0.5 + 1i):   4,
	(-0.5 + 1.5i): 2,
	(-0.5 + 2i):   1,
	(-1 - 2i):     1,
	(-1 - 1.5i):   2,
	(-1 - 1i):     3,
	(-1 - 0.5i):   5,
	(-1 + 0i):     255,
	(-1 + 0.5i):   5,
	(-1 + 1i):     3,
	(-1 + 1.5i):   2,
	(-1 + 2i):     1,
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Test the ComputeIterations method with a battery of numbers and known
// answers.
func TestComputeIterations(t *testing.T) {
	for n, solution := range iterationTests {
		iterations := ComputeIterations(n)
		if iterations != solution {
			t.Errorf("%v: expected %d, obtained %d", n, solution, iterations)
		}
	}
}
