package main

import (
	"fmt"
	"math"
)

type Coeff struct {
	cos float64
	sin float64
}

const SAMPLE_SIZE = 8
const PI = math.Pi

func main() {
	var samples [SAMPLE_SIZE]float64
	var coefficients [SAMPLE_SIZE]Coeff
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
}

func IDFT(coefficients []Coeff, recovered []float64) {
	for r_idx := range recovered {
		real := 0.0
		for f_idx := range coefficients {
			ratio := float64(r_idx) / float64(len(recovered))
			angle := 2.0 * PI * float64(f_idx) * ratio
			cos_A, sin_B := coefficients[f_idx].cos, coefficients[f_idx].sin

			real += cos_A * math.Cos(angle)
			real -= sin_B * math.Sin(angle)
		}
		real /= float64(len(recovered))
		recovered[r_idx] = real
	}
}

func meanSquared(expected []float64, actual []float64) float64 {
	error := 0.0
	for i := range expected {
		diff := expected[i] - actual[i]
		error += diff * diff
	}
	return error
}

