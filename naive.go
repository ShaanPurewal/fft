package main

import "math"

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
}

