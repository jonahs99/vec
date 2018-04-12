// Package vec contains a 2d vector type and function for doing math with them
package vec

// Vec is a 2d vector type
type Vec struct {
	X, Y float64
}

// Set copies from another Vec
func (v *Vec) Set(v1 Vec) *Vec {
	v.X, v.Y = v1.X, v1.Y
	return v
}

// Add adds the components from another Vec
func (v *Vec) Add(v1 Vec) *Vec {
	v.X += v1.X
	v.Y += v1.Y
	return v
}

// Add returns the sum of 2 vectors
func Add(v1, v2 Vec) Vec {
	return Vec{v1.X + v2.X, v1.Y + v2.Y}
}

// Times scalar multiplies
func (v *Vec) Times(a float64) *Vec {
	v.X *= a
	v.Y *= a
	return v
}

// Times returns a scaled vector
func Times(v Vec, a float64) Vec {
	return Vec{v.X * a, v.Y * a}
}

// Sub subtracts the components from another Vec
func (v *Vec) Sub(v1 Vec) *Vec {
	v.X -= v1.X
	v.Y -= v1.Y
	return v
}

// Sub returns the difference of 2 vectors
func Sub(v1, v2 Vec) Vec {
	return Vec{v1.X - v2.X, v1.Y - v2.Y}
}

// Div scalar divides
func (v *Vec) Div(a float64) *Vec {
	v.X /= a
	v.Y /= a
	return v
}

// Div returns the scalar division of a vector
func Div(v1 Vec, a float64) Vec {
	return Vec{v1.X * a, v1.Y * a}
}
