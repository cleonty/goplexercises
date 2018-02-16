package main

import (
	"fmt"
	"unicode"
)

func main() {
	months := [...]string{1: "Январь", 2: "Февраль", 3: "Март", 4: "Апрель", 5: "Май", 6: "Июнь", 7: "Июль", 8: "Август", 9: "Сентябрь", 10: "Октябрь", 11: "Ноябрь", 12: "Декабрь"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	for _, s := range Q2 {
		for _, q := range summer {
			if s == q {
				fmt.Printf("%s находится в обоих срезах\n", s)
			}
		}
	}
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	b := []int{0, 1, 2, 3, 4, 5}
	reverse(b[:2])
	reverse(b[2:])
	reverse(b)
	fmt.Println(b)
	{
		a := make([]int, 3, 10)
		fmt.Printf("%T %d %d %d\n", a, a, len(a), cap(a))
	}
	{
		a := []int{0, 1, 2, 3, 4}
		b := appendInt(a, 5)
		fmt.Println(b, cap(b))
		b = appendInt2(b, b...)
		fmt.Println(b, cap(b))
	}
	{
		data := []string{"one", "", "three"}
		fmt.Println(data)
		fmt.Println(nonempty(data))
		fmt.Println(data)
	}
	{
		data := []string{"one", "", "three"}
		fmt.Println(data)
		fmt.Println(nonempty2(data))
		fmt.Println(data)
	}
	{
		var stack []string
		stack = append(stack, "leonty")
		top := stack[len(stack)-1]
		fmt.Println(top, stack)
		stack = stack[:len(stack)-1]
		fmt.Println(stack)
	}
	{
		x := []int{1, 2, 3, 4, 5, 6, 7, 8}
		fmt.Println(remove(x, 2))
	}
	{
		x := []int{1, 2, 3, 4, 5, 6, 7, 8}
		fmt.Println(remove2(x, 2))
	}
	{
		x := []int{1, 2, 3, 4, 4, 4, 5, 6, 6, 6, 6, 7, 8, 8, 8, 9}
		fmt.Println(removeDuplicates(x))
		x = []int{}
		fmt.Println(removeDuplicates(x))
		x = nil
		fmt.Println(removeDuplicates(x))
	}
	{
		s := "a   b   c     d"
		fmt.Println(string(squeezeSpaces([]byte(s))))
	}
	{
		s := "a   b   c     d"
		fmt.Println(string(reverseBytes([]byte(s))))
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	var zlen = len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	var zlen = len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func removeDuplicates(slice []int) []int {
	for i := 0; i < len(slice)-1; {
		if slice[i] == slice[i+1] {
			slice = remove(slice, i+1)
		} else {
			i++
		}
	}
	return slice
}

func removeByte(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func squeezeSpaces(slice []byte) []byte {
	for i := 0; i < len(slice)-1; {
		if unicode.IsSpace(rune(slice[i])) && unicode.IsSpace(rune(slice[i+1])) {
			slice = removeByte(slice, i+1)
		} else {
			i++
		}
	}
	return slice
}

func reverseBytes(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
