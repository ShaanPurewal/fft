package internal

import "math"

var EVALS = []func(float64) float64{
	Constant,
	Sin,
	Cos,
	Custom,
}

func Constant(_ float64) float64 { return 420.6767 }

func Sin(ratio float64) float64 {
	return 4.67 + 3 * math.Sin(2.0 * PI * ratio)
}

func Cos(ratio float64) float64 {
	return 4.67 + 3 * math.Cos(2.0 * PI * ratio)
}

func Custom(ratio float64) float64 {
	return 100 + 7 * math.Sin(2.0 * PI * ratio) + 4 * math.Exp(ratio)
}

func GenerateSamples(samples []float64, fn func(float64) float64) {
	for i := range samples {
		samples[i] = fn(float64(i) / float64(len(samples)))
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
