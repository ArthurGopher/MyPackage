package mymath

import "math"

func Abs(x float64) float64 {
		
	return math.Abs(x)
}

func Max(a, b float64) float64 {
	return math.Max(a,b)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n,x)
}
