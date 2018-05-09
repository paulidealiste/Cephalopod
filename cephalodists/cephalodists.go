// Package cephalodists provides probability distribution functions
package cephalodists

import "math"

// ProbabiltyF returns the value of F probability distribution function for the supplied value and degrees of freedom
func ProbabiltyF(x float64, d1 float64, d2 float64) float64 {
	denominator := x * B(d1/2, d2/2)
	numerator := math.Sqrt((math.Pow(d1*x, d1) * math.Pow(d2, d2)) / (math.Pow(d1*x+d2, d1+d2)))
	return numerator / denominator
}

// B represents Beta function
func B(x float64, y float64) float64 {
	return math.Gamma(x) * math.Gamma(y) / math.Gamma(x+y)
}
