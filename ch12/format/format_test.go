package format

import (
	"fmt"
	"time"
)

func ExampleAny() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
    fmt.Println(Any(x))
    fmt.Println(Any(d))
    fmt.Println(Any([]int64{x})) // "[]int64 0x8202b87b0"
    fmt.Println(Any([]time.Duration{d})) // "[]time.Duration    
}
