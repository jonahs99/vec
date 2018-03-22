package vec

import "testing"

func TestIntersect(t *testing.T) {
	cases := []struct {
		o1, d1, o2, d2 Vec
		t1, t2         float64
		result         int
	}{
		{Vec{0, 0}, Vec{1, 0}, Vec{3, 2}, Vec{0, -1},
			3, 2, 1},
		{Vec{1, 0}, Vec{1, 1}, Vec{1, 2}, Vec{1, -1},
			1, 1, 1},
		{Vec{0, 0}, Vec{1, 0}, Vec{2, 0}, Vec{-1, 0},
			0, 0, -1},
		{Vec{0, 0}, Vec{1, 0}, Vec{2, 0.1}, Vec{-1, 0},
			0, 0, 0},
	}

	for _, c := range cases {
		t1, t2, result := Intersect(c.o1, c.d1, c.o2, c.d2)
		if t1 != c.t1 || t2 != c.t2 || result != c.result {
			t.Errorf("intersect(%v,%v,%v,%v) == %v,%v,%v, want %v,%v,%v",
				c.o1, c.d1, c.o2, c.d2, t1, t2, result, c.t1, c.t2, c.result)
		}
	}
}
