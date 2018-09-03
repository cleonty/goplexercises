package popcount

import (
	"sync"
)

var pcOnce sync.Once
var pc [256]byte

func initPC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	pcOnce.Do(initPC)
	var result int
	for i := uint8(0); i < 8; i++ {
		result += int(pc[byte(x>>i*8)])
	}
	return result
}
