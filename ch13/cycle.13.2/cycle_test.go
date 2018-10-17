package cycle

import (
	"bytes"
	"testing"
)

func TestCycle(t *testing.T) {
	one := 1

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	type mystring string

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{"foo", false},
		// slices
		{[]string{"foo"}, false},
		// slice cycles
		{cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{},
			false,
		},
		// pointers
		{&one, false},
		{new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
	} {
		if HasCycle(test.x) != test.want {
			t.Errorf("HasCycle(%#v) = %t",
				test.x, !test.want)
		}
	}
}
