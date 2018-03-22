package vec

// Intersect calculates the intersection point of two lines
func Intersect(o1, d1, o2, d2 Vec) (float64, float64, int) {
	o := *o2.Sub(o1)
	det := Cross(d1, d2)

	if det == 0 {
		if Cross(o, d1) == 0 {
			return 0, 0, -1
		}
		return 0, 0, 0
	}

	t1 := (o.X*d2.Y - o.Y*d2.X) / det
	t2 := (-o.X*d1.Y + o.Y*d1.X) / det

	return t1, -t2, 1
}
