package main

import (
	"fmt"
	"runtime"
	"reflect"
	"strings"
	"os"
	"strconv"
	"github.com/shaanpurewal/fft/internal"
)

/*

   NOTE: This implementation assumes that the input samples are real-valued,
   this is exploited when taking advantage of coefficient symmetries.

   Discrete Fourier Transform (w/ Inverse)
     - Naive
     - FFT
       - RADIX-2 (w/ naive fallback)
       - TODO: MIXED-RADIX
       - TODO: Bluestien
       - TODO: In-Place
     - TODO: Parallel (Optimal)

   (Shaan Purewal)
   
*/

const USAGE = "Expected: ./fft <SAMPLE_SIZE> <IMPLEMENTATION>"

func main() {
	// Parse arguments
	args := os.Args[1:]

	DFT := internal.DFFT
	IDFT := internal.IDFFT

	sample_size := 64 * 1024
	if len(args) > 0 {
		val, err := strconv.Atoi(args[0])
		if err != nil { panic(USAGE) }
		sample_size = val

		if len(args) < 2 { panic(USAGE) }
		switch args[1] {
		case "naive":
			DFT = internal.DFT
			IDFT = internal.IDFT
		case "fast":
			DFT = internal.DFFT
			IDFT = internal.IDFFT
		default:
			DFT = internal.DFFT
			IDFT = internal.IDFFT
		}
	}

	nyquist := sample_size / 2 + 1

	for _, fn := range internal.EVALS {
		eval(fn, DFT, IDFT, sample_size, nyquist)
	}
}

func eval(fn func(float64) float64, DFT func([]float64, []internal.Coeff), IDFT func([]internal.Coeff, []float64), sample_size int, nyquist int) {
	samples := make([]float64, sample_size)
	coefficients := make([]internal.Coeff, nyquist)
	recovered := make([]float64, sample_size)

	eval_name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	eval_name = eval_name[strings.LastIndex(eval_name, ".")+1:]
	fmt.Printf("\n%s: \n", eval_name)
	
	// Aquire (gen)
	internal.GenerateSamples(samples[:], fn)
	fmt.Printf(" - %d samples\n", sample_size)
	// Perform DFT
	DFT(samples[:], coefficients[:])
	fmt.Println(" - Finished FT")

	// Perform IDFT
	IDFT(coefficients[:], recovered[:])
	fmt.Println(" - Finished IFT")

	// Compare Results
	fmt.Printf(" -> MSE: %.30f\n", internal.MeanSquared(samples[:], recovered[:]))
}
