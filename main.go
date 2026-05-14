package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/*

   NOTE: This implementation assumes that the input samples are real-valued,
   this is exploited when taking advantage of coefficient symmetries.

   NAIVE Discrete Fourier Transform (w/ Inverse)
   (Shaan Purewal)
   
*/

type Coeff struct {
	cos float64
	sin float64
}

const (
	SAMPLE_SIZE = 10_000
	PI = math.Pi
	NYQUIST = SAMPLE_SIZE / 2 + 1
)

func main() {
	var samples [SAMPLE_SIZE]float64
	var coefficients [NYQUIST]Coeff
	var recovered [SAMPLE_SIZE]float64

	// Aquire (gen)
	generateSamples(samples[:], custom)
	
	// Perform DFT
	DFT(samples[:], coefficients[:])

	// Perform IDFT
	IDFT(coefficients[:], recovered[:])

	// Compare Results
	fmt.Printf("\nMSE: %.30f\n", meanSquared(samples[:], recovered[:]))
}

func constant(_ int) float64 { return 10 }
func linear(i int) float64 { return 10 + float64(i) }
func sin(i int) float64 {
	ratio := float64(i) / float64(SAMPLE_SIZE)
	return math.Sin(2.0 * PI * ratio)
}
func exp(i int) float64 {
	ratio := float64(i) / float64(SAMPLE_SIZE)
	return math.Exp(ratio)
}
func custom(i int) float64 { return 100 + sin(i) + exp(i) }

func generateSamples(samples []float64, fn func(int) float64) {
	for i := range samples {
		samples[i] = fn(i)
	}
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	fmt.Printf("\n%d samples generated w/ '%s'\n", len(samples), name)
}

func DFT(samples []float64, coefficients []Coeff) {
	for f_idx := range coefficients {
		coeff := Coeff{}
		for s_idx := range samples {
			ratio := float64(s_idx) / float64(len(samples))
			angle := 2.0 * PI * float64(f_idx) * ratio
			
			coeff.cos += samples[s_idx] * math.Cos(angle)
			coeff.sin -= samples[s_idx] * math.Sin(angle)
		}
		coefficients[f_idx] = coeff
	}
	fmt.Println("Finished performing Fourier Transform")
}

func IDFT(coefficients []Coeff, recovered []float64) {
	N := len(recovered)
	for r_idx := range recovered {
		real := 0.0
		for f_idx := range coefficients {
			ratio := float64(r_idx) / float64(N)
			angle := 2.0 * PI * float64(f_idx) * ratio
			cos_A, sin_B := coefficients[f_idx].cos, coefficients[f_idx].sin

			term := cos_A * math.Cos(angle) - sin_B * math.Sin(angle)
			evenNyquist := (N % 2 == 0) && (f_idx == N/2)
			
			if f_idx > 0 && !evenNyquist{
				term *= 2
			}
			real += term
		}
		real /= float64(N) // Normalize
		recovered[r_idx] = real
	}
	fmt.Println("Finished performing Inverse Fourier Transform")
}

func meanSquared(expected []float64, actual []float64) float64 {
	error := 0.0
	for i := range expected {
		diff := expected[i] - actual[i]
		error += diff * diff
	}
	return error
}

