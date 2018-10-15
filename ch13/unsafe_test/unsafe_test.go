package unsafe_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Alignof(false))
}

func TestOffsetof(t *testing.T) {
	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(x.a), unsafe.Sizeof(x.b), unsafe.Sizeof(x.c))
	fmt.Println(unsafe.Alignof(x), unsafe.Alignof(x.a), unsafe.Alignof(x.b), unsafe.Alignof(x.c))
	fmt.Println(unsafe.Offsetof(x.a), unsafe.Offsetof(x.b), unsafe.Offsetof(x.c))
}

func float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
func TestUnsafePointer(t *testing.T) {
	fmt.Printf("%#016x\n", float64bits(1.0))
}

func TestArithmetic(t *testing.T) {
	var x struct {
		a bool
		b int16
		c []int
	}
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
}
