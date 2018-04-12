package vec

import "math"

// New returns a copy of a Vec
func New(v Vec) Vec {
	return Vec{v.X, v.Y}
}

// NewXY returns a fresh Vec with given coordinates
func NewXY(x, y float64) Vec {
	return Vec{x, y}
}

// NewPolar returns a fresh Vec from polar coordinates
func NewPolar(r, t float64) Vec {
	return Vec{r * math.Cos(t), r * math.Sin(t)}
}
