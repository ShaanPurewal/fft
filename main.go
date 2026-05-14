package main

import (
	"fmt"
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

const (
	SAMPLE_SIZE = 1_048_576
	NYQUIST = SAMPLE_SIZE / 2 + 1
)

func main() {
	var samples [SAMPLE_SIZE]float64
	var coefficients [NYQUIST]internal.Coeff
	var recovered [SAMPLE_SIZE]float64

	// Aquire (gen)
	internal.GenerateSamples(samples[:], internal.Custom)
	fmt.Printf("\n%d samples generated\n", SAMPLE_SIZE)
	
	// Perform DFT
	internal.DFFT(samples[:], coefficients[:])
	fmt.Println("Finished performing Fourier Transform")

	// Perform IDFT
	internal.IDFFT(coefficients[:], recovered[:])
	fmt.Println("Finished performing Inverse Fourier Transform")

	// Compare Results
	fmt.Printf("\nMSE: %.30f\n", internal.MeanSquared(samples[:], recovered[:]))
}

