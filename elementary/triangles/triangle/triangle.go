package triangle

import "math"

type Triangle struct {
	Name  string
	Sides [3]float64
	Area  float64
}

func New(name string, sides [3]float64) Triangle {
	t := Triangle{name, sides, getArea(sides)}
	return t
}

func getArea(sides [3]float64) float64 {
	var s, p, a float64

	for _, side := range sides {
		s += side
	}

	p = s / 2
	a = p

	for _, side := range sides {
		a *= (p - side)
	}

	return math.Sqrt(a)
}
