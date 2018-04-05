package intset

import (
	"testing"
)

func Test(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	expectedX := "{1 9 144}"
	expectedLenX := 3
	if x.String() != expectedX {
		t.Errorf("expected %s, got %s\n", expectedX, x.String())
	}
	if x.Len() != expectedLenX {
		t.Errorf("expected %d, got %d\n", expectedLenX, x.Len())
	}
	y.Add(9)
	y.Add(42)
	y.Add(58)
	y.Add(99)
	y.Add(512)
	y.Add(1024)
	y.Remove(1024)
	expectedY := "{9 42 58 99 512}"
	if y.String() != expectedY {
		t.Errorf("expected %s, got %s\n", expectedY, y.String())
	}
	expectedLenY := 5
	if y.Len() != expectedLenY {
		t.Errorf("expected %d, got %d\n", expectedLenY, y.Len())
	}

	x.UnionWith(&y)
	expectedUnion := "{1 9 42 58 99 144 512}"
	if x.String() != expectedUnion {
		t.Errorf("expected %s, got %s\n", expectedUnion, x.String())
	}
	expectedUnionLen := 7
	if x.Len() != expectedUnionLen {
		t.Errorf("expected %d, got %d\n", expectedUnionLen, x.Len())
	}
	if !x.Has(1) {
		t.Errorf("expected x has 1, x is %s\n", x.String())
	}
	if !x.Has(9) {
		t.Errorf("expected x has 9, x is %s\n", x.String())
	}
	if !x.Has(42) {
		t.Errorf("expected x has 42, x is %s\n", x.String())
	}
	if !x.Has(144) {
		t.Errorf("expected x has 144, x is %s\n", x.String())
	}
	z := x.Copy()
	if z.String() != expectedUnion {
		t.Errorf("expected %s, got %s\n", expectedUnion, z.String())
	}
	if z.Len() != expectedUnionLen {
		t.Errorf("expected %d, got %d\n", expectedUnionLen, z.Len())
	}
	x.Clear()
	if x.Len() != 0 {
		t.Errorf("expected len=0, got %d", x.Len())
	}
	if x.String() != "{}" {
		t.Errorf("expected {}, got %s\n", x.String())
	}

}

func TestIntersection(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(9)
	x.Add(144)

	y.Add(2)
	y.Add(9)
	y.Add(144)

	x.IntersectWith(&y)
	expected := "{9 144}"
	if x.String() != expected {
		t.Errorf("expected %s, got %s", expected, x.String())
	}
}
func TestDifference(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(9)
	x.Add(144)
	x.Add(20000)

	y.Add(2)
	y.Add(9)
	y.Add(144)

	x.DifferenceWith(&y)
	expected := "{1 20000}"
	if x.String() != expected {
		t.Errorf("expected %s, got %s", expected, x.String())
	}
}
func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(9)
	x.Add(144)
	x.Add(20000)

	y.Add(2)
	y.Add(9)
	y.Add(144)
	y.Add(1440000)

	x.SymmetricDifferenceWith(&y)
	expected := "{1 2 20000 1440000}"
	if x.String() != expected {
		t.Errorf("expected %s, got %s", expected, x.String())
	}
}

func TestElems(t *testing.T) {
	var x IntSet
	expectedElements := []uint{1, 9, 144, 2000}
	for _, elem := range expectedElements {
		x.Add(int(elem))
	}
	elements := x.Elems()
	if len(elements) != len(expectedElements) {
		t.Errorf("len of slices doesn't match, got %d, expected %d\n", len(elements), len(expectedElements))
	}
	for i, elem := range elements {
		if elem != expectedElements[i] {
			t.Errorf("Element [%d] doesn't match, got %d, expected %d\n", i, elem, expectedElements[i])
		}
	}

}
