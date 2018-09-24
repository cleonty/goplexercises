package reflect_test

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func ExampleReflect() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	// Output:
	// int
    // *os.File
    // 3
    // 3
    // <int Value>
}
