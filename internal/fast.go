package internal

import "math"

/*
   Implementation: RADIX-2, w/ naive DFT fallback
*/

func DFFT(samples []float64, coefficients []Coeff) {
	X := recFFT(samples)

	// Only need N/2 coeffs (X will have the full N)
	for i := range coefficients {
		coefficients[i] = X[i]
	}
}

func recFFT(samples []float64) []Coeff {
	// Base Case: Non-Even (kinda slow)
	N := len(samples)
	if N % 2 == 1 {
		coefficients := make([]Coeff, N)
		DFT(samples, coefficients)
		return coefficients
	}

	// Recusive: split into even and odd freq bins
	even := make([]float64, N/2)
	odd := make([]float64, N/2)

	for i := 0; i < N/2; i++ {
		even[i] = samples[2*i]
		odd[i] = samples[2*i + 1]
	}

	E := recFFT(even)
	O := recFFT(odd)

	coefficients := make([]Coeff, N)
	
	for f_idx := 0; f_idx < N/2; f_idx++ {
		angle := -2.0 * PI * float64(f_idx) / float64(N)

		c := math.Cos(angle)
		s := math.Sin(angle)

		tcos := O[f_idx].cos*c - O[f_idx].sin*s
		tsin := O[f_idx].cos*s + O[f_idx].sin*c

		coefficients[f_idx].cos = E[f_idx].cos + tcos
		coefficients[f_idx].sin = E[f_idx].sin + tsin

		coefficients[f_idx + N/2].cos = E[f_idx].cos - tcos
		coefficients[f_idx + N/2].sin = E[f_idx].sin - tsin
	}

	return coefficients
}

func IDFFT(coefficients []Coeff, recovered []float64) {
	N := len(recovered)

	// Rebuild conjugate-symmetric half
	X := make([]Coeff, N)
	copy(X, coefficients)

	for k := 1; k < len(coefficients) && k < N-k; k++ {
		X[N-k] = Coeff{
			cos: X[k].cos,
			sin: -X[k].sin,
		}
	}

	x := recIFFT(X)

	// Normalize
	for i := range recovered {
		recovered[i] = x[i].cos / float64(N)
	}
}

func recIFFT(coefficients []Coeff) []Coeff {
	// Base Case: Non-Even (kinda slow)
	N := len(coefficients)
	if N % 2 == 1 {
		recovered := make([]float64, N)
		IDFT(coefficients, recovered)

		out := make([]Coeff, N)
		for i, v := range recovered {
			out[i].cos = v
		}
		return out
	}

	even := make([]Coeff, N/2)
	odd := make([]Coeff, N/2)

	for i := 0; i < N/2; i++ {
		even[i] = coefficients[2*i]
		odd[i] = coefficients[2*i + 1]
	}

	E := recIFFT(even)
	O := recIFFT(odd)

	recovered := make([]Coeff, N)

	for f_idx := 0; f_idx < N/2; f_idx++ {
		angle := 2.0 * PI * float64(f_idx) / float64(N)

		c := math.Cos(angle)
		s := math.Sin(angle)

		tcos := O[f_idx].cos*c - O[f_idx].sin*s
		tsin := O[f_idx].cos*s + O[f_idx].sin*c

		recovered[f_idx].cos = E[f_idx].cos + tcos
		recovered[f_idx].sin = E[f_idx].sin + tsin

		recovered[f_idx + N/2].cos = E[f_idx].cos - tcos
		recovered[f_idx + N/2].sin = E[f_idx].sin - tsin
	}

	return recovered
}

