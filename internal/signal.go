package internal

import "math"

func Custom(i int, size int) float64 {
	ratio := float64(i) / float64(size)
	return 100 + 7 * math.Sin(2.0 * PI * ratio) + 4 * math.Exp(ratio)
}

func GenerateSamples(samples []float64, fn func(int, int) float64) {
	for i := range samples {
		samples[i] = fn(i, len(samples))
	}
}

func MeanSquared(expected []float64, actual []float64) float64 {
	error := 0.0
	for i := range expected {
		diff := expected[i] - actual[i]
		error += diff * diff
	}
	return error
}
