package main

import "math"

func custom(i int) float64 {
	ratio := float64(i) / float64(SAMPLE_SIZE)
	return 100 + 7 * math.Sin(2.0 * PI * ratio) + 4 * math.Exp(ratio)
}

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
