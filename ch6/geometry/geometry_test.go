package geometry

import "testing"

func TestDistance1(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	d := p.Distance(q)
	if d != 5 {
		t.Error("Expected 5, got ", d)
	}
}
func TestDistance2(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	d := Distance(p, q)
	if d != 5 {
		t.Error("Expected 5, got ", d)
	}
}
func TestPathDistance(t *testing.T) {
	perim := Path{
		{1, 1}, {5, 1},
		{5, 4},
		{1, 1},
	}
	d := perim.Distance()
	if d != 12 {
		t.Error("Expected 12, got ", d)
	}
}

func TestScaleBy(t *testing.T) {
	p := Point{1, 2}
	q := Point{2, 4}
	p.ScaleBy(2)
	if p != q {
		t.Error("Expected {2, 4}, got ", p)
	}
}
