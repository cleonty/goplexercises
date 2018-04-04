package main

import (
	"fmt"
	"math"
	"time"
)

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

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(3)
	scaleP(5)
	fmt.Println(p)

	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() { r.Launch() })
	time.AfterFunc(10*time.Second, r.Launch)

	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 7)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)

	path := Path{{1, 2}, {3, 4}}
	path.TranslateBy(Point{1, 2}, false)
	fmt.Println(path)
}

type Rocket struct{}

func (r *Rocket) Launch() {}
