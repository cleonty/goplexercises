package cycle

import (
	"reflect"
	"unsafe"
)

func hasCycle(x reflect.Value, seen map[comparision]bool) bool {
	if !x.IsValid() {
		return false
	}
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparision{xptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Bool:
		return false
	case reflect.String:
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return false
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return false
	case reflect.Float32, reflect.Float64:
		return false
	case reflect.Complex64, reflect.Complex128:
		return false
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return false
	case reflect.Ptr, reflect.Interface:
		return hasCycle(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if hasCycle(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if hasCycle(x.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if hasCycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	panic("unreachable")
}

type comparision struct {
	x unsafe.Pointer
	t reflect.Type
}

// HasCycle сообщает есть ли цикл в х.
func HasCycle(x interface{}) bool {
	seen := make(map[comparision]bool)
	return hasCycle(reflect.ValueOf(x), seen)
}
