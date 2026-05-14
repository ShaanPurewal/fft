package main

import (
	"fmt"
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

func main() {
	// Parse arguments
	args := os.Args[1:]

	sample_size := 64 * 1024
	if len(args) > 0 {
		val, err := strconv.Atoi(args[0])
		if err != nil { panic("Expected: ./fft SAMPLE_SIZE") }
		sample_size = val
	}

	nyquist := sample_size / 2 + 1
	
	samples := make([]float64, sample_size)
	coefficients := make([]internal.Coeff, nyquist)
	recovered := make([]float64, sample_size)

	// Aquire (gen)
	internal.GenerateSamples(samples[:], internal.Custom)
	fmt.Printf("\n%d samples generated\n", sample_size)
	
	// Perform DFT
	internal.DFFT(samples[:], coefficients[:])
	fmt.Println("Finished performing Fourier Transform")

	// Perform IDFT
	internal.IDFFT(coefficients[:], recovered[:])
	fmt.Println("Finished performing Inverse Fourier Transform")

	// Compare Results
	fmt.Printf("\nMSE: %.30f\n", internal.MeanSquared(samples[:], recovered[:]))
}

