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

// Times scalar multiplies
func (v *Vec) Times(a float64) *Vec {
	v.X *= a
	v.Y *= a
	return v
}

// Sub subtracts the components from another Vec
func (v *Vec) Sub(v1 Vec) *Vec {
	v.X -= v1.X
	v.Y -= v1.Y
	return v
}

// Div scalar divides
func (v *Vec) Div(a float64) *Vec {
	v.X /= a
	v.Y /= a
	return v
}
