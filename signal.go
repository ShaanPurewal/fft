package main

import "math"

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

func meanSquared(expected []float64, actual []float64) float64 {
	error := 0.0
	for i := range expected {
		diff := expected[i] - actual[i]
		error += diff * diff
	}
	return error
}
