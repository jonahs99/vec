package vec

import (
	"math"
)

var zero = Vec{0, 0}

// Zero returns the zero vector
func Zero() Vec {
	return zero
}

// Mag2 calculates the squared magnitude
func Mag2(v Vec) float64 {
	return math.Pow(v.X, 2) + math.Pow(v.Y, 2)
}

// Mag calculates the magnitude
func Mag(v Vec) float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Theta gets the polar angle coordinate
func Theta(v Vec) float64 {
	return math.Atan2(v.Y, v.X)
}

// Dot calculates the inner product
func Dot(v1 Vec, v2 Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Cross calculates the magnitude of the R3 cross product
func Cross(v1 Vec, v2 Vec) float64 {
	return v1.X*v2.Y - v2.X*v1.Y
}

// Dist2 calculates the magnitude of the difference vector
func Dist2(v1, v2 Vec) float64 {
	return math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2)
}

// Dist calculates the magnitude of the difference vector
func Dist(v1, v2 Vec) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}

// Norm returns a counter-clockwise points normal
func Norm(v Vec) Vec {
	return Vec{-v.Y, v.X}
}

// Reflect returns v1 reflected across n (assumed to be a UNIT vector!)
func Reflect(v1, n Vec) Vec {
	// r = v1 - 2(v1 . n)n
	n.Times(2 * Dot(v1, n))
	v1.Sub(n)
	return v1
}

// Truncate returns a vector of at maximum v
func Truncate(v Vec, a float64) Vec {
	mag := Mag(v)
	if mag > a {
		return Times(v, a/mag)
	}
	return v
}
