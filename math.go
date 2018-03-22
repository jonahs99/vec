package vec

import (
	"math"
)

// Mag2 calculates the squared magnitude
func Mag2(v Vec) float64 {
	return math.Pow(v.X, 2) + math.Pow(v.Y, 2)
}

// Mag calculates the magnitude
func Mag(v Vec) float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Dot calculates the inner product
func Dot(v1 Vec, v2 Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Cross calculates the magnitude of the R3 cross product
func Cross(v1 Vec, v2 Vec) float64 {
	return v1.X*v2.Y - v2.X*v1.Y
}
