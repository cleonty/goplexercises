package geometry

import "math"

// Point represent 2d Point
type Point struct{ X, Y float64 }

// Path represents a path
type Path []Point

// Distance does
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance does
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy scales point by a factor
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Distance returns path's distance
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
