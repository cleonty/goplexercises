package intlist

import "testing"

func testSumWithSingleElementList(t *testing.T) {
	list := &IntList{1, nil}
	sum := list.Sum()
	if sum != 1 {
		t.Error("Expected 1, got ", sum)
	}
}

func testSumWithTwoElementsInList(t *testing.T) {
	list := &IntList{1, &IntList{2, nil}}
	sum := list.Sum()
	if sum != 3 {
		t.Error("Expected 3, got ", sum)
	}
}

func testSumWithNilList(t *testing.T) {
	var list *IntList
	sum := list.Sum()
	if sum != 0 {
		t.Error("Expected 0, got ", sum)
	}
}
