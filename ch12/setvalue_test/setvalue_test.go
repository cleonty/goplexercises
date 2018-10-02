package setvalue_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestSetValue(t *testing.T) {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)

	d.Set(reflect.ValueOf(5))
	fmt.Println(x)

	t.Run("test invalid type", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		d.Set(reflect.ValueOf(int64(55)))
	})

	t.Run("test non-addressable value", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		х := 2
		b := reflect.ValueOf(х)
		b.Set(reflect.ValueOf(3))
	})

}

func TestSetvalue2(t *testing.T) {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)               // OK, x = 2
	rx.Set(reflect.ValueOf(3)) // OK, x = 3
	//rx.SetString("hello")            // Аварийная ситуация: string не присваиваемо int
	//rx.Set(reflect.ValueOf("hello")) // Аварийная ситуация: string не присваиваемо int
	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	//ry.SetInt(2) // Аварийная ситуация:
	// Setlnt вызван для интерфейса Value
	ry.Set(reflect.ValueOf(3)) // OK, у = int(3)
	//ry.SetString("hello")      // Аварийная ситуация: SetString
	// вызван для интерфейса Value
	ry.Set(reflect.ValueOf("hello")) // OK, у = "hello"
}

func TestStdout(t *testing.T) {
	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())
}
